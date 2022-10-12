package handler

import (
	"campfirereads/internal/domain"
	"github.com/gin-gonic/gin"
)

type Search interface {
	Search(c *gin.Context)
}

func (a *APIHandler) Search(c *gin.Context) {
	res, err := a.google.GetBooks(domain.BookSearch{Query: c.Query("query")})
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(200, res)
}

func (a *APIHandler) ISBN(c *gin.Context) {
	res, err := a.google.GetISBN(c.Param("isbn"))
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	c.JSON(200, res)
}