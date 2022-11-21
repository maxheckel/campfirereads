package handler

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/maxheckel/campfirereads/internal/domain"
	"github.com/maxheckel/campfirereads/internal/service/payments"
	"io"
	"math/rand"
	"time"
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

func RandStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
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

	url, err := a.payments.CheckoutURL(booksWithListings, RandStringRunes(7))
	if err != nil {
		var priceMismatchErr *payments.PriceMismatchErr
		if errors.As(err, &priceMismatchErr) {
			c.JSON(400, gin.H{
				"type":  "price_mismatch",
				"error": err.Error(),
				"data": map[string]interface{}{
					"isbn":        priceMismatchErr.ISBN,
					"listingType": priceMismatchErr.ListingType,
					"actualPrice": priceMismatchErr.ActualPriceInCents,
				},
			})
			return
		}

		var outOfStockErr *payments.OutOfStockErr
		if errors.As(err, &outOfStockErr) {
			c.JSON(400, gin.H{
				"type":  "out_of_stock",
				"error": err.Error(),
				"data": map[string]interface{}{
					"isbn":        outOfStockErr.ISBN,
					"listingType": outOfStockErr.ListingType,
				},
			})
			return
		}

		// Default error
		c.JSON(500, gin.H{"error": err})
		return
	}
	c.JSON(200, gin.H{"url": url})
}

func (a *APIHandler) Receipt(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(500, gin.H{"error": errors.New("missing url parameter ID")})
	}
	res, err := a.payments.GetOrder(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	c.JSON(200, res)

}
