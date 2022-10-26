package service

import (
	"github.com/maxheckel/campfirereads/internal/config"
	"github.com/maxheckel/campfirereads/internal/domain"
)

type NYT interface {
	GetBestSellers() ([]*domain.BestSeller, error)
}

type nyt struct {
	cfg config.Config
}

