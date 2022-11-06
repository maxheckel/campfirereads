package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/maxheckel/campfirereads/internal/domain"
	"strings"
	"sync"
)

type BestSellers interface {
	GetBestSellers(c *gin.Context)
	Popular(c *gin.Context)
}

type BestSellerResponse struct {
	Lists []NYTListWithGoogleBooks `json:"lists"`
}

type NYTListWithGoogleBooks struct {
	List  *domain.List   `json:"list"`
	Books []*domain.Book `json:"books"`
}

var (
	popularISBNs = []string{
		"9780525463467", // my side of the mountain
		"9780385056199", // where the red fern grows
		"9780330351690", // into the wild
		"9780060115456", // old yeller
		"9780007136599", // the fellowship of the ring
		"9780691014647", // walden
		"9780061233326", // pilgrim at tinker creek
	}
)

func (a *APIHandler) GetBestSellers(c *gin.Context) {
	cacheKey := "bestsellers-resp"
	cache, err := a.cache.Read(cacheKey)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	if cachedResponseBytes, ok := cache.([]byte); ok {
		res := BestSellerResponse{}
		err = json.Unmarshal(cachedResponseBytes, &res)
		if err != nil {
			c.JSON(500, gin.H{"error": err})
			return
		}
		c.JSON(200, res)
		return
	}

	// If the last time it was stored was today
	bestSellerNYT, err := a.nyt.GetBestSellers()
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	ISBNList := []string{}
	for _, list := range bestSellerNYT.Results.Lists {
		if strings.Contains(list.ListNameEncoded, "-e-book") || strings.Contains(list.ListNameEncoded, "audio") {
			continue
		}
		if len(list.Books) == 0 {
			continue
		}
		for _, book := range list.Books {
			if len(book.PrimaryIsbn13) > 0 {
				ISBNList = append(ISBNList, book.PrimaryIsbn10)
			}
		}
	}
	books := a.booksFromISBNs(ISBNList)
	isbnToBook := map[string]*domain.Book{}
	for _, book := range books {
		if book == nil {
			continue
		}
		if book.VolumeInfo == nil {
			continue
		}
		for _, isbn := range book.VolumeInfo.IndustryIdentifiers {
			if isbn.Type == "ISBN_13" {
				isbnToBook[isbn.Identifier] = book
			}
		}
	}

	res := BestSellerResponse{}
	for _, list := range bestSellerNYT.Results.Lists {
		if strings.Contains(list.ListNameEncoded, "-e-book") || strings.Contains(list.ListNameEncoded, "audio") {
			continue
		}
		if len(list.Books) == 0 {
			continue
		}
		list := list
		resList := NYTListWithGoogleBooks{}
		resList.List = &list
		for _, book := range list.Books {
			if isbnToBook[book.PrimaryIsbn13] == nil {
				continue
			}
			resList.Books = append(resList.Books, isbnToBook[book.PrimaryIsbn13])
		}
		resList.List.Books = nil
		res.Lists = append(res.Lists, resList)
	}
	err = a.cache.Write(cacheKey, res, 24*60*60)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	c.JSON(200, res)
}

func (a *APIHandler) booksFromISBNs(ISBNList []string) []*domain.Book {
	books := make([]*domain.Book, len(ISBNList))
	wg := sync.WaitGroup{}
	for i, isbn := range ISBNList {
		i := i
		wg.Add(1)
		go func(isbn string) {
			defer wg.Done()
			book, err := a.google.GetISBN(isbn, 0)
			if err != nil {
				panic(err)
			}
			books[i] = book
		}(isbn)
	}
	wg.Wait()
	return books
}

func (a *APIHandler) Category(c *gin.Context) {
	cat := strings.TrimSpace(c.Param("category"))
	if cat == "" {
		c.JSON(500, gin.H{"error": "You must provide a category"})
		return
	}
	if cat == "popular" {
		books := a.booksFromISBNs(popularISBNs)
		c.JSON(200, NYTListWithGoogleBooks{
			List: &domain.List{
				ListID:          0,
				ListName:        "",
				ListNameEncoded: "",
				DisplayName:     "Popular",
				Updated:         "",
				ListImage:       nil,
				ListImageWidth:  nil,
				ListImageHeight: nil,
				Books:           nil,
			},
			Books: books,
		})
		return
	}

	res, err := a.nyt.GetCategory(cat)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	ISBNs := []string{}
	for _, book := range res.Results.Books {
		ISBNs = append(ISBNs, book.PrimaryIsbn13)
	}
	books := a.booksFromISBNs(ISBNs)

	response := &NYTListWithGoogleBooks{List: &res.Results, Books: books}
	c.JSON(200, response)
}

func (a *APIHandler) Popular(c *gin.Context) {
	books := a.booksFromISBNs(popularISBNs)
	c.JSON(200, books)
}
