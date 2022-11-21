package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/maxheckel/campfirereads/internal/config"
	"github.com/maxheckel/campfirereads/internal/domain"
	"github.com/maxheckel/campfirereads/internal/service/cache"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Google interface {
	GetBooks(search domain.BookSearch) (*domain.BookSearchResult, error)
	GetISBN(isbn string, sleep int) (*domain.Book, error)
	GetVolumeByID(id, isbn string, sleep int) (*domain.Book, error)
}

type google struct {
	config *config.Config
	cache  cache.Service
}

func NewGoogle(cfg *config.Config, cache cache.Service) Google {
	return &google{config: cfg, cache: cache}
}

func (g google) GetVolumeByID(id, isbn string, sleep int) (*domain.Book, error) {
	// If it's not numeric then we should just skip
	if _, err := strconv.Atoi(isbn); err != nil {
		return nil, nil
	}
	if isbn == "" {
		return nil, errors.New("you must provide a valid isbn")
	}
	cacheKey := isbn
	cachedBook, err := g.cache.Read(cacheKey)
	if err != nil {
		return nil, err
	}
	if b, ok := cachedBook.(*domain.Book); ok {
		return b, nil
	}
	if id == "" {
		return nil, errors.New("you must provide a valid volume id")
	}
	res := &domain.Book{}
	reqURL := fmt.Sprintf("https://www.googleapis.com/books/v1/volumes/%s?key=%s", id, g.config.GoogleAPIKey)
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
	if response.StatusCode != 200 {
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
	err = g.cache.Write(cacheKey, res, 24*60*60)
	return res, err
}

func (g google) GetISBN(isbn string, sleep int) (*domain.Book, error) {
	// If it's not numeric then we should just skip
	if _, err := strconv.Atoi(isbn); err != nil {
		return nil, nil
	}
	cacheKey := isbn
	cachedBook, err := g.cache.Read(cacheKey)
	if err != nil {
		return nil, err
	}
	if b, ok := cachedBook.(*domain.Book); ok {
		return b, nil
	}
	res := &domain.BookSearchResult{}
	reqURL := fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?key=%s&q=isbn:%s", g.config.GoogleAPIKey, isbn)
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
		ISBN := book.ISBN()
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
