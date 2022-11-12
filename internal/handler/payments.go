package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/maxheckel/campfirereads/internal/domain"
	"io"
)

type PaymentHandler interface {
	GetPublicKey(c *gin.Context)
}

func (a *APIHandler) GetPublicKey(c *gin.Context) {
	key, err := a.payments.GetPublicKey()
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	c.JSON(200, gin.H{"key": key})
}

func (a *APIHandler) GetCheckoutURL(c *gin.Context) {
	var booksWithListings []*domain.BookWithListing
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	err = json.Unmarshal(body, &booksWithListings)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	url, err := a.payments.GetCheckoutURL(booksWithListings)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	c.JSON(200, gin.H{"url": url})
}
