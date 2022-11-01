package service

import (
	"encoding/json"
	"fmt"
	"github.com/maxheckel/campfirereads/internal/config"
	"github.com/maxheckel/campfirereads/internal/domain"
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
}

func NewGoogle(cfg *config.Config) Google {
	return &google{config: cfg}
}

func (g google) GetISBN(isbn string, sleep int) (*domain.Book, error) {
	res := &domain.BookSearchResult{}

	query, err := url.Parse(fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?langRestrict=en&q=isbn:%s", isbn))
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
	return &res.Items[0], err
}

func (g google) GetBooks(search domain.BookSearch) (*domain.BookSearchResult, error) {
	res := &domain.BookSearchResult{}

	query, err := url.Parse(fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?langRestrict=en&q=%s", url.QueryEscape(search.Query)))
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
