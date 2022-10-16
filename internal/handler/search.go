package handler

import (
	"campfirereads/internal/domain"
	"github.com/gin-gonic/gin"
	"sync"
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
	res, err := a.amazon.ISBNToListings(c.Param("isbn"))
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	wg := sync.WaitGroup{}
	for index := range res {
		wg.Add(1)
		go func(index int) {
			err := a.amazon.ListingToPriceInCents(res[index])
			if err != nil {
				panic(err)
			}
			wg.Done()
		}(index)
	}
	wg.Wait()

	c.JSON(200, res)
}
