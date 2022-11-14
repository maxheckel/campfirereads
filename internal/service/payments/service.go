package payments

import (
	"github.com/maxheckel/campfirereads/internal/domain"
)

type PriceMismatchErr struct {
	ISBN               string
	ListingType        string
	Name               string
	ActualPriceInCents int64
}

type OutOfStockErr struct {
	ISBN        string
	ListingType string
	Name        string
}

func (e *OutOfStockErr) Error() string {
	return "Some items are longer in stock, please check your cart and checkout again"
}

func (e *PriceMismatchErr) Error() string {
	return "Some items in your carts price have changed during your session, their new price is updated in your cart.  Please confirm this the new prices then checkout again."
}

type Service interface {
	CreateOrRetrieveProduct(book domain.Book) (id string, err error)
	GetCheckoutURL([]*domain.BookWithListing) (id string, err error)
	GetReceipt(id string) (*domain.Receipt, error)
	GetPublicKey() (string, error)
}
