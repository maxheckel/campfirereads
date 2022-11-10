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
	cache  cache.Service
}

func NewGoogle(cfg *config.Config, cache cache.Service) Google {
	return &google{config: cfg, cache: cache}
}

func (g google) GetISBN(isbn string, sleep int) (*domain.Book, error) {
	cacheKey := isbn
	cachedBook, err := g.cache.Read(cacheKey)
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
			return g.GetISBN(isbn, sleep+(sleep*2))
		}

		return nil, err
	}
	if err != nil {
		return nil, err
	}
	b := &res.Items[0]
	err = g.cache.Write(cacheKey, b, 24*60*60)
	return b, err
}

func (g google) GetBooks(search domain.BookSearch) (*domain.BookSearchResult, error) {
	res := &domain.BookSearchResult{}
	query, err := url.Parse(fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?maxResults=20&startIndex=%d&langRestrict=en&q=%s", search.Offset, url.QueryEscape(search.Query)))
	if err != nil {
		return nil, err
	}
	fmt.Println(query.String())
	response, err := http.Get(query.String())
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		search.Backoff = search.Backoff + 1
		if search.Backoff <= 5 {
			time.Sleep(time.Duration(search.Backoff) * time.Second)
			fmt.Println(fmt.Sprintf("Backoff for book search, waiting seconds: %d", search.Backoff))
			return g.GetBooks(search)
		}
		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)

		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("google error: %s", string(body))
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, err
	}
	for _, book := range res.Items {
		ISBN := ""
		for _, identifier := range book.VolumeInfo.IndustryIdentifiers {
			if identifier.Type == "ISBN_13" {
				ISBN = identifier.Identifier
				break
			}
		}
		if ISBN == "" {
			continue
		}
		err = g.cache.Write(ISBN, book, 24*60*60)
		if err != nil {
			return nil, err
		}
	}
	return res, err
}
