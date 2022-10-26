package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/maxheckel/campfirereads/internal/domain"
	"sync"
)

type BestSellers interface {
	GetBestSellers(c *gin.Context)
}

func (a *APIHandler) GetBestSellers(c *gin.Context) {
	bestSellerNYT, err := a.nyt.GetBestSellers()
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	ISBNList := []string{}
	for _, bestSeller := range bestSellerNYT {
		if len(bestSeller.RanksHistory) > 0 {
			ISBNList = append(ISBNList, bestSeller.RanksHistory[0].PrimaryIsbn13)
			continue
		}
		if len(bestSeller.Isbns) > 0 {
			ISBNList = append(ISBNList, bestSeller.Isbns[0].Isbn13)
			continue
		}
	}
	//
	//bookSearch := make(chan *domain.Book, len(ISBNList))
	//errChan := make(chan error)
	books := make([]*domain.Book, len(ISBNList))
	wg := sync.WaitGroup{}
	for i, isbn := range ISBNList {
		i := i
		wg.Add(1)
		go func(isbn string) {
			defer wg.Done()
			book, err := a.google.GetISBN(isbn)
			if err != nil {
				panic(err)
			}
			books[i] = book
		}(isbn)
	}
	wg.Wait()
	c.JSON(200, books)

}
