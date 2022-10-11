package handler

import (
	"github.com/gin-gonic/gin"
)

type Search interface {
	Search(c *gin.Context)
}

func (a *APIHandler) Search(c *gin.Context) {
	c.JSON(200, gin.H{"success": true})
}
