package payments

import (
	"github.com/maxheckel/campfirereads/internal/domain"
)

type Service interface {
	CreateOrRetrieveProduct(book domain.Book) (id string, err error)
	CreateCharge([]*domain.BookWithListing) (id string, err error)
	GetPublicKey() (string, error)
}
