package payments

import (
	"errors"
	"fmt"
	"github.com/maxheckel/campfirereads/internal/config"
	"github.com/maxheckel/campfirereads/internal/domain"
	"github.com/maxheckel/campfirereads/internal/service"
	"github.com/stripe/stripe-go/v73"
	"github.com/stripe/stripe-go/v73/checkout/session"
	"github.com/stripe/stripe-go/v73/price"
	"github.com/stripe/stripe-go/v73/product"
	"strings"
	"sync"
)

type stripeService struct {
	private     string
	public      string
	frontendURL string
	amazon      service.Amazon
}

const isbnKey = "isbn"
const amazonURLKey = "amazon_url"

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
func (s *stripeService) GetCheckoutURL(books []*domain.BookWithListing) (string, error) {
	lineItems := make([]*stripe.CheckoutSessionLineItemParams, len(books))
	errs := make([]error, len(books))
	wg := sync.WaitGroup{}
	for i, b := range books {
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
			stripePrice, err := s.createPrice(stripeProduct, b.Listing)
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
		SuccessURL:         stripe.String(fmt.Sprintf("%s/checkout/success", s.frontendURL)),
		CancelURL:          stripe.String(fmt.Sprintf("%s/cart", s.frontendURL)),
		LineItems:          lineItems,
		Mode:               stripe.String(string(stripe.CheckoutSessionModePayment)),
		PaymentMethodTypes: []*string{stripe.String("card")},
		ShippingAddressCollection: &stripe.CheckoutSessionShippingAddressCollectionParams{
			AllowedCountries: []*string{stripe.String("US")},
		},
		//ShippingOptions: []*stripe.CheckoutSessionShippingOptionParams{
		//	{
		//		ShippingRateData: &stripe.CheckoutSessionShippingOptionShippingRateDataParams{
		//			FixedAmount: &stripe.CheckoutSessionShippingOptionShippingRateDataFixedAmountParams{
		//				Amount:   stripe.Int64(0),
		//				Currency: stripe.String(string(stripe.CurrencyUSD)),
		//			},
		//			DisplayName: stripe.String("Free shipping"),
		//			DeliveryEstimate: &stripe.CheckoutSessionShippingOptionShippingRateDataDeliveryEstimateParams{
		//				Minimum: &stripe.CheckoutSessionShippingOptionShippingRateDataDeliveryEstimateMinimumParams{
		//					Unit:  stripe.String("business_day"),
		//					Value: stripe.Int64(5),
		//				},
		//				Maximum: &stripe.CheckoutSessionShippingOptionShippingRateDataDeliveryEstimateMaximumParams{
		//					Unit:  stripe.String("business_day"),
		//					Value: stripe.Int64(7),
		//				},
		//			},
		//		},
		//	},
		//},
	}
	stripeSession, err := session.New(params)
	if err != nil {
		return "", err
	}
	return stripeSession.URL, nil
}

func (s *stripeService) createPrice(product *stripe.Product, listing *domain.AmazonListing) (*stripe.Price, error) {
	unitAmount := stripe.Int64(int64(listing.PriceInCents + 1000))
	query := &stripe.PriceSearchParams{}
	query.Query = *stripe.String(fmt.Sprintf("product:'%s'", product.ID))
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
	return price.New(params)
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
			found = true
			b.Listing = l
			break
		}
	}
	if !found {
		return errors.New("could not find exact listing for book")
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
	description := book.Book.VolumeInfo.Description
	words := strings.Split(description, " ")
	description = strings.Join(words[0:25], " ") + "..."
	createReq := &stripe.ProductParams{}
	createReq.Name = &book.Book.VolumeInfo.Title
	createReq.Images = append(createReq.Images, &book.Book.VolumeInfo.ImageLinks.Thumbnail)
	createReq.Description = &description
	createReq.AddMetadata(amazonURLKey, book.Listing.URL.Path)
	createReq.AddMetadata(isbnKey, book.Book.ISBN())
	return product.New(createReq)
}

func (s *stripeService) GetPublicKey() (string, error) {
	return s.public, nil
}
