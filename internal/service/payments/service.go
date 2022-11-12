package payments

import (
	"github.com/maxheckel/campfirereads/internal/domain"
)

type Service interface {
	CreateOrRetrieveProduct(book domain.Book) (id string, err error)
	GetCheckoutURL([]*domain.BookWithListing) (id string, err error)
	GetReceipt(id string) (*domain.Receipt, error)
	GetPublicKey() (string, error)
}
