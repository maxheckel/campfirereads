package payments

import (
	"github.com/maxheckel/campfirereads/internal/config"
	"github.com/maxheckel/campfirereads/internal/domain"
	"github.com/stripe/stripe-go/v73"
)

type stripeService struct {
	private string
	public  string
}

func Stripe(cfg *config.Config) Service {
	stripe.Key = cfg.StripePrivateAPIKey
	return &stripeService{private: cfg.StripePrivateAPIKey, public: cfg.StripePublicAPIKey}
}

func (s *stripeService) CreateOrRetrieveProduct(book domain.Book) (id string, err error) {
	
}
func (s *stripeService) CreateCharge(books []*domain.BookWithListing) (id string, err error) {
	panic("blah")
}

func (s *stripeService) GetPublicKey() (string, error) {
	return s.public, nil
}
