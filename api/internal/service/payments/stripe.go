package payments

import (
	"fmt"
	"github.com/maxheckel/campfirereads/internal/config"
	"github.com/maxheckel/campfirereads/internal/domain"
	"github.com/maxheckel/campfirereads/internal/service"
	"github.com/stripe/stripe-go/v73"
	"github.com/stripe/stripe-go/v73/checkout/session"
	"github.com/stripe/stripe-go/v73/price"
	"github.com/stripe/stripe-go/v73/product"
	"net/url"
	"strings"
	"sync"
	"time"
)

type stripeService struct {
	private     string
	public      string
	frontendURL string
	amazon      service.Amazon
}

const isbnKey = "ISBN"
const amazonURLKey = "amazon_url"
const listingTypeKey = "listing_type"
const authorsKey = "authors"

func Stripe(cfg *config.Config, amazon service.Amazon) Service {
	stripe.Key = cfg.StripePrivateAPIKey
	return &stripeService{
		private:     cfg.StripePrivateAPIKey,
		public:      cfg.StripePublicAPIKey,
		amazon:      amazon,
		frontendURL: cfg.FrontendHost,
	}
}

func (s *stripeService) CreateOrRetrieveProduct(book domain.Book) (id string, err error) {
	panic("blah")
}
func (s *stripeService) CheckoutURL(booksWithListings []*domain.BookWithListing, internalOrderID string) (string, error) {
	lineItems := make([]*stripe.CheckoutSessionLineItemParams, len(booksWithListings))
	errs := make([]error, len(booksWithListings))
	wg := sync.WaitGroup{}
	for i, b := range booksWithListings {
		i := i
		wg.Add(1)
		go func(b *domain.BookWithListing) {
			defer wg.Done()
			stripeProduct, err := s.productForBook(b)
			if err != nil {
				errs[i] = err
				return
			}
			err = s.latestListingPrice(err, b)
			if err != nil {
				errs[i] = err
				return
			}
			stripePrice, err := s.createOrRetrievePrice(stripeProduct, b.Listing)
			if err != nil {
				errs[i] = err
				return
			}
			lineItems[i] = &stripe.CheckoutSessionLineItemParams{
				Price:    &stripePrice.ID,
				Quantity: stripe.Int64(1),
			}
		}(b)
	}
	wg.Wait()
	for _, err := range errs {
		if err != nil {
			return "", err
		}
	}
	params := &stripe.CheckoutSessionParams{
		SuccessURL:         stripe.String(fmt.Sprintf("%s/receipt/{CHECKOUT_SESSION_ID}?clearCart=true", s.frontendURL)),
		CancelURL:          stripe.String(fmt.Sprintf("%s/cart", s.frontendURL)),
		LineItems:          lineItems,
		ClientReferenceID:  stripe.String(internalOrderID),
		Mode:               stripe.String(string(stripe.CheckoutSessionModePayment)),
		PaymentMethodTypes: []*string{stripe.String("card")},
		ShippingAddressCollection: &stripe.CheckoutSessionShippingAddressCollectionParams{
			AllowedCountries: []*string{stripe.String("US")},
		},
		ShippingOptions: []*stripe.CheckoutSessionShippingOptionParams{
			{
				ShippingRateData: &stripe.CheckoutSessionShippingOptionShippingRateDataParams{
					Type: stripe.String("fixed_amount"),
					FixedAmount: &stripe.CheckoutSessionShippingOptionShippingRateDataFixedAmountParams{
						Amount:   stripe.Int64(499),
						Currency: stripe.String(string(stripe.CurrencyUSD)),
					},
					DisplayName: stripe.String("Standard Shipping"),
					DeliveryEstimate: &stripe.CheckoutSessionShippingOptionShippingRateDataDeliveryEstimateParams{
						Minimum: &stripe.CheckoutSessionShippingOptionShippingRateDataDeliveryEstimateMinimumParams{
							Unit:  stripe.String("business_day"),
							Value: stripe.Int64(5),
						},
						Maximum: &stripe.CheckoutSessionShippingOptionShippingRateDataDeliveryEstimateMaximumParams{
							Unit:  stripe.String("business_day"),
							Value: stripe.Int64(7),
						},
					},
				},
			},
		},
	}
	stripeSession, err := session.New(params)
	if err != nil {
		return "", err
	}
	return stripeSession.URL, nil
}

func (s *stripeService) createOrRetrievePrice(product *stripe.Product, listing *domain.AmazonListing) (*stripe.Price, error) {
	unitAmount := stripe.Int64(listing.PriceInCents + config.SmokeCostPerUnit)
	if *unitAmount <= config.SmokeCostPerUnit {
		return nil, fmt.Errorf("unit amount is too low: %d", unitAmount)
	}
	query := &stripe.PriceSearchParams{}
	query.Query = *stripe.String(fmt.Sprintf("product:'%s' AND metadata['%s']:'%s'", product.ID, listingTypeKey, listing.Type))
	iter := price.Search(query)
	for iter.Next() {
		if iter.Price().UnitAmount == *unitAmount {
			return iter.Price(), nil
		}
	}

	params := &stripe.PriceParams{
		Product:    stripe.String(product.ID),
		Currency:   stripe.String(string(stripe.CurrencyUSD)),
		UnitAmount: unitAmount,
	}
	params.AddMetadata(listingTypeKey, listing.Type)
	return price.New(params)
}

func (s *stripeService) GetOrder(id string) (*domain.Receipt, error) {
	receipt := &domain.Receipt{}
	sessionParams := stripe.CheckoutSessionParams{}
	sessionParams.AddExpand("customer")
	sessionParams.AddExpand("payment_intent")
	sessionParams.AddExpand("shipping_cost.shipping_rate")
	stripeSession, err := session.Get(id, &sessionParams)
	if err != nil {
		return nil, err
	}

	receipt.OrderedOn = time.Unix(stripeSession.Created, 0)

	receipt.ShippingCost = domain.ShippingCost{
		AmountInCents: stripeSession.ShippingCost.AmountTotal,
		MinDays:       stripeSession.ShippingCost.ShippingRate.DeliveryEstimate.Minimum.Value,
		MaxDays:       stripeSession.ShippingCost.ShippingRate.DeliveryEstimate.Maximum.Value,
		Name:          stripeSession.ShippingCost.ShippingRate.DisplayName,
	}

	receipt.Customer = domain.Customer{
		Name:        stripeSession.CustomerDetails.Name,
		PhoneNumber: stripeSession.CustomerDetails.Phone,
		Email:       stripeSession.CustomerDetails.Email,
	}

	receipt.Shipping = domain.Address{
		Street1: stripeSession.ShippingDetails.Address.Line1,
		Street2: stripeSession.ShippingDetails.Address.Line2,
		City:    stripeSession.ShippingDetails.Address.City,
		State:   stripeSession.ShippingDetails.Address.State,
		Zip:     stripeSession.ShippingDetails.Address.PostalCode,
	}

	receipt.Billing = domain.Address{
		Street1: stripeSession.PaymentIntent.Charges.Data[0].BillingDetails.Address.Line1,
		Street2: stripeSession.PaymentIntent.Charges.Data[0].BillingDetails.Address.Line2,
		City:    stripeSession.PaymentIntent.Charges.Data[0].BillingDetails.Address.City,
		State:   stripeSession.PaymentIntent.Charges.Data[0].BillingDetails.Address.State,
		Zip:     stripeSession.PaymentIntent.Charges.Data[0].BillingDetails.Address.PostalCode,
	}

	paymentDetails := stripeSession.PaymentIntent.Charges.Data[0].PaymentMethodDetails
	receipt.PaymentType = string(paymentDetails.Type)
	if paymentDetails.Type == "card" {
		receipt.PaymentIdentifier = fmt.Sprintf("Last 4: %s, Expires: %d/%d", paymentDetails.Card.Last4, paymentDetails.Card.ExpMonth, paymentDetails.Card.ExpYear)
	}
	receipt.TotalInCents = stripeSession.AmountTotal
	receipt.OrderID = stripeSession.ClientReferenceID

	params := &stripe.CheckoutSessionListLineItemsParams{}
	params.Session = stripe.String(id)
	params.AddExpand("data.price.product")
	i := session.ListLineItems(params)
	for i.Next() {
		li := i.LineItem()
		amazonURL, err := url.Parse(fmt.Sprintf("https://amazon.com%s", li.Price.Product.Metadata[amazonURLKey]))
		if err != nil {
			return nil, err
		}
		if li != nil {

			receipt.Books = append(receipt.Books, domain.BookWithListing{
				Book: &domain.Book{
					VolumeInfo: &domain.VolumeInfo{
						Title:       li.Price.Product.Name,
						Description: li.Price.Product.Description,
						ImageLinks:  &domain.Images{SmallThumbnail: li.Price.Product.Images[0], Thumbnail: li.Price.Product.Images[0]},
						IndustryIdentifiers: []domain.Identifier{
							{
								Type:       "ISBN_13",
								Identifier: li.Price.Product.Metadata[isbnKey],
							},
						},
						Authors: []string{
							li.Price.Product.Metadata[authorsKey],
						},
					},
				},
				Listing: &domain.AmazonListing{
					PriceInCents: li.Price.UnitAmount,
					URL:          amazonURL,
					Type:         li.Price.Metadata[listingTypeKey],
					ISBN:         li.Price.Product.Metadata[isbnKey],
				},
			})
		}
	}

	return receipt, nil
}

func (s *stripeService) latestListingPrice(err error, b *domain.BookWithListing) error {
	// Fetch a fresh price
	listings, err := s.amazon.ISBNToPrices(b.Book.ISBN())
	if err != nil {
		return err
	}
	found := false
	for _, l := range listings {
		if l.URL.Path == b.Listing.URL.Path && l.Type == b.Listing.Type {
			if b.Listing.PriceInCents != l.PriceInCents && l.PriceInCents > 0 {
				return &PriceMismatchErr{
					ISBN:               b.Book.ISBN(),
					ActualPriceInCents: l.PriceInCents,
					Name:               b.Book.VolumeInfo.Title,
					ListingType:        b.Listing.Type,
				}
			}
			if l.PriceInCents <= 0 {
				return &OutOfStockErr{ISBN: b.Book.ISBN(), Name: b.Book.VolumeInfo.Title, ListingType: b.Listing.Type}
			}
			found = true
			b.Listing = l
			break
		}
	}
	if !found {
		return &OutOfStockErr{ISBN: b.Book.ISBN(), Name: b.Book.VolumeInfo.Title, ListingType: b.Listing.Type}
	}
	if b.Listing.PriceInCents <= 0 {
		return fmt.Errorf("no price found for ISBN %s", b.Book.ISBN())
	}
	return nil
}

func (s *stripeService) productForBook(book *domain.BookWithListing) (*stripe.Product, error) {
	params := &stripe.ProductSearchParams{}
	params.Query = *stripe.String(fmt.Sprintf("metadata['%s']:'%s' AND metadata['%s']:'%s'", isbnKey, amazonURLKey, book.Book.ISBN(), book.Listing.URL.Path))
	iter := product.Search(params)
	for iter.Next() {
		return iter.Product(), nil
	}
	createReq := &stripe.ProductParams{}

	description := book.Book.VolumeInfo.Description
	if description != "" {
		words := strings.Split(description, " ")
		if len(words) > 25 {
			description = strings.Join(words[0:25], " ") + "..."
		}
		createReq.Description = &description
	}

	createReq.Name = &book.Book.VolumeInfo.Title
	createReq.Images = append(createReq.Images, &book.Book.VolumeInfo.ImageLinks.Thumbnail)

	createReq.AddMetadata(amazonURLKey, book.Listing.URL.Path)
	createReq.AddMetadata(isbnKey, book.Book.ISBN())
	createReq.AddMetadata(authorsKey, strings.Trim(strings.Join(book.Book.VolumeInfo.Authors, ", "), ", "))
	return product.New(createReq)
}

func (s *stripeService) GetPublicKey() (string, error) {
	return s.public, nil
}
