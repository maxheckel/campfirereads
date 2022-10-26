package server

import (
	"github.com/gin-gonic/gin"
	"github.com/maxheckel/campfirereads/internal/handler"
	"github.com/maxheckel/campfirereads/internal/service"
)

func NewAPI() (*App, error) {
	srv, err := New("API")
	if err != nil {
		panic(err)
	}
	h := handler.NewAPI(service.NewGoogle(srv.Config), service.NewAmazon(), service.NewNYT(srv.Config))
	srv.Gin.GET("/search", h.Search)
	srv.Gin.GET("/isbn/:isbn", h.ISBN)
	srv.Gin.GET("/bestsellers", h.GetBestSellers)
	// Healthcheck
	srv.Gin.GET("/", func(context *gin.Context) {
		context.Writer.Write([]byte("OK"))
	})
	return srv, nil
}
