package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/maxheckel/campfirereads/internal/domain"
	"sync"
)

type Search interface {
	Search(c *gin.Context)
}

type GetBookResponse struct {
	Book     *domain.Book            `json:"book"`
	Listings []*domain.AmazonListing `json:"listings"`
}

func (a *APIHandler) Search(c *gin.Context) {
	res, err := a.google.GetBooks(domain.BookSearch{Query: c.Query("query")})
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(200, res.Items)
}

func (a *APIHandler) ISBN(c *gin.Context) {
	ISBN := c.Param("isbn")
	var amazonListings []*domain.AmazonListing
	var book *domain.Book
	var err error
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup, isbn string, listings []*domain.AmazonListing) {
		defer wg.Done()
		amazonListings, err = a.amazon.ISBNToPrices(isbn)

	}(&wg, ISBN, amazonListings)
	wg.Add(1)
	go func(wg *sync.WaitGroup, isbn string) {
		defer wg.Done()

		book, err = a.google.GetISBN(isbn, 0)

	}(&wg, ISBN)
	wg.Wait()
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	response := GetBookResponse{
		Book:     book,
		Listings: amazonListings,
	}

	c.JSON(200, response)
}
