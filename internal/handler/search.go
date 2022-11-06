package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/maxheckel/campfirereads/internal/domain"
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
	book, err := a.google.GetISBN(ISBN, 0)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	response := GetBookResponse{
		Book: book,
	}

	c.JSON(200, response)
}
func (a *APIHandler) Price(c *gin.Context) {
	ISBN := c.Param("isbn")
	prices, err := a.amazon.ISBNToPrices(ISBN)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	response := GetBookResponse{
		Listings: prices,
	}

	c.JSON(200, response)
}
