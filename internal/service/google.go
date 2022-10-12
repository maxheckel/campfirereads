package service

import (
	"campfirereads/internal/config"
	"campfirereads/internal/domain"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Google interface {
	GetBooks(search domain.BookSearch) (*domain.BookSearchResult, error)
	GetISBN(isbn string) (*domain.SearchResult, error)
}

type google struct {
	config *config.Config
}

func NewGoogle(cfg *config.Config) Google {
	return &google{config: cfg}
}

func (g google) GetISBN(isbn string) (*domain.SearchResult, error) {
	res := &domain.SearchResult{}

	query, err := url.Parse(fmt.Sprintf("https://www.googleapis.com/customsearch/v1?key=%s&lr=lang_en&cx=5002d862df2554bc7&q=%s", g.config.GoogleAPIKey, isbn))
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
	return res, err
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
