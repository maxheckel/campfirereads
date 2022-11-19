package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/maxheckel/campfirereads/internal/domain"
	"strconv"
)

type Search interface {
	Search(c *gin.Context)
}

type GetBookResponse struct {
	Book     *domain.Book            `json:"book"`
	Listings []*domain.AmazonListing `json:"listings"`
}

func (a *APIHandler) Search(c *gin.Context) {
	if c.Query("query") == "" {
		c.JSON(200, domain.BookSearchResult{})
		return
	}
	page := c.Query("page")
	pageInt, err := strconv.Atoi(page)
	if page != "" && err != nil {
		pageInt = 0
	}

	if pageInt <= 0 {
		pageInt = 1
	}
	// page - 1 because the first page has no offset
	offset := 20 * (pageInt - 1)
	res, err := a.google.GetBooks(domain.BookSearch{Query: c.Query("query"), Offset: offset})
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	c.JSON(200, res)
}

func (a *APIHandler) ISBN(c *gin.Context) {
	ISBN := c.Param("isbn")
	vID := c.Query("v")
	var book *domain.Book
	var err error
	if vID == "" {
		book, err = a.google.GetISBN(ISBN, 0)
	} else {
		book, err = a.google.GetVolumeByID(vID, ISBN, 0)
	}
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
