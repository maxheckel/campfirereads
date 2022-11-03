package service

import (
	"encoding/json"
	"fmt"
	"github.com/maxheckel/campfirereads/internal/config"
	"github.com/maxheckel/campfirereads/internal/domain"
	"github.com/maxheckel/campfirereads/internal/service/cache"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Google interface {
	GetBooks(search domain.BookSearch) (*domain.BookSearchResult, error)
	GetISBN(isbn string, sleep int) (*domain.Book, error)
}

type google struct {
	config *config.Config
	cache  cache.Cache
}

func NewGoogle(cfg *config.Config, cache cache.Cache) Google {
	return &google{config: cfg, cache: cache}
}

func (g google) GetISBN(isbn string, sleep int) (*domain.Book, error) {
	cachedBook, err := g.cache.Read(isbn)
	if err != nil {
		return nil, err
	}
	if b, ok := cachedBook.(*domain.Book); ok {
		return b, nil
	}
	res := &domain.BookSearchResult{}
	reqURL := fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?langRestrict=en&key=%s&q=isbn:%s", g.config.GoogleAPIKey, isbn)
	query, err := url.Parse(reqURL)
	if err != nil {
		return nil, err
	}
	response, err := http.Get(query.String())
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, res)
	if len(res.Items) == 0 {
		if response.StatusCode == 429 && strings.Contains(string(body), "RATE_LIMIT_EXCEEDED") {
			if sleep >= 5 {
				return nil, err
			}
			fmt.Printf("Backoff on ISBN %s, waiting %d seconds\n", isbn, sleep)
			time.Sleep(time.Second * time.Duration(sleep))
			return g.GetISBN(isbn, sleep+1)
		}

		return nil, err
	}
	if err != nil {
		return nil, err
	}
	b := &res.Items[0]
	err = g.cache.Write(isbn, b, 24*60*60)
	return b, err
}

func (g google) GetBooks(search domain.BookSearch) (*domain.BookSearchResult, error) {
	res := &domain.BookSearchResult{}

	query, err := url.Parse(fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?key=%s&langRestrict=en&q=%s", g.config.GoogleAPIKey, url.QueryEscape(search.Query)))
	if err != nil {
		return nil, err
	}
	fmt.Println(query.String())
	response, err := http.Get(query.String())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, res)
	return res, err
}
