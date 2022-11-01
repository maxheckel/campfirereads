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
	srv.Gin.Use(CORSMiddleware())
	srv.Gin.GET("/search", h.Search)
	srv.Gin.GET("/isbn/:isbn", h.ISBN)
	srv.Gin.GET("/bestsellers", h.GetBestSellers)
	srv.Gin.GET("/popular", h.Popular)
	// Healthcheck
	srv.Gin.GET("/", func(context *gin.Context) {
		context.Writer.Write([]byte("OK"))
	})
	return srv, nil
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
