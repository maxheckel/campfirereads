package service

import (
	"encoding/json"
	"fmt"
	"github.com/maxheckel/campfirereads/internal/config"
	"github.com/maxheckel/campfirereads/internal/domain"
	"io"
	"net/http"
)

type NYT interface {
	GetBestSellers() (*domain.AllListsBestSellers, error)
	GetCategory(category string) (*domain.GetBestSellerList, error)
}

type nyt struct {
	cfg *config.Config
}

func NewNYT(cfg *config.Config) NYT {
	return &nyt{cfg: cfg}
}

func (n *nyt) GetCategory(category string) (*domain.GetBestSellerList, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.nytimes.com/svc/books/v3/lists/current/%s.json?api-key=%s", category, n.cfg.NYTAPIKey))
	if err != nil {
		return nil, err
	}
	res := &domain.GetBestSellerList{}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (n *nyt) GetBestSellers() (*domain.AllListsBestSellers, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.nytimes.com/svc/books/v3/lists/full-overview.json?api-key=%s", n.cfg.NYTAPIKey))
	if err != nil {
		return nil, err
	}
	res := &domain.AllListsBestSellers{}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
