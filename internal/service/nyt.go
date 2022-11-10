package service

import (
	"encoding/json"
	"fmt"
	"github.com/maxheckel/campfirereads/internal/config"
	"github.com/maxheckel/campfirereads/internal/domain"
	"github.com/maxheckel/campfirereads/internal/service/cache"
	"io"
	"net/http"
)

type NYT interface {
	GetBestSellers() (*domain.AllListsBestSellers, error)
	GetCategory(category string) (*domain.GetBestSellerList, error)
}

type nyt struct {
	cfg   *config.Config
	cache cache.Service
}

func NewNYT(cfg *config.Config, cache cache.Service) NYT {
	return &nyt{cfg: cfg, cache: cache}
}

func (n *nyt) GetCategory(category string) (*domain.GetBestSellerList, error) {
	cacheKey := fmt.Sprintf("category-%s", category)
	categoryCache, err := n.cache.Read(cacheKey)
	if err != nil {
		return nil, err
	}
	if b, ok := categoryCache.(*domain.GetBestSellerList); ok {
		return b, nil
	}
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
	n.cache.Write(cacheKey, res, 24*60*60)
	return res, nil
}

func (n *nyt) GetBestSellers() (*domain.AllListsBestSellers, error) {
	bestSellersCache, err := n.cache.Read("bestsellers")
	if err != nil {
		return nil, err
	}
	if b, ok := bestSellersCache.(*domain.AllListsBestSellers); ok {
		return b, nil
	}
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
	n.cache.Write("bestsellers", res, 24*60*60)
	return res, nil
}
