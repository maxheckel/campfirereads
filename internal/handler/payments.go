package handler

import (
	"github.com/gin-gonic/gin"
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
