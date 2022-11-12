package server

import (
	"github.com/gin-gonic/gin"
	"github.com/maxheckel/campfirereads/internal/handler"
	"github.com/maxheckel/campfirereads/internal/service"
	"github.com/maxheckel/campfirereads/internal/service/cache"
	"github.com/maxheckel/campfirereads/internal/service/payments"
)

func NewAPI() (*App, error) {
	srv, err := New("API")
	if err != nil {
		panic(err)
	}
	var cacheService cache.Service
	switch srv.Config.CacheDriver {
	case "memory":
		cacheService = &cache.Memory{}
	case "memcache-local":
		cacheService = cache.NewMemcache(srv.Config.CacheAddress)
	case "memcache-appengine":
		cacheService = cache.NewAppEngineMemcache()
	default:
		cacheService = &cache.Memory{}

	}
	amazon := service.NewAmazon(cacheService)
	merchant := payments.Stripe(srv.Config, amazon)
	h := handler.NewAPI(service.NewGoogle(srv.Config, cacheService), amazon, service.NewNYT(srv.Config, cacheService), cacheService, merchant)
	srv.Gin.Use(CORSMiddleware())
	srv.Gin.GET("/search", h.Search)
	srv.Gin.GET("/isbn/:isbn", h.ISBN)
	srv.Gin.GET("/isbn/:isbn/price", h.Price)
	srv.Gin.GET("/category/:category", h.Category)
	srv.Gin.GET("/bestsellers", h.GetBestSellers)
	srv.Gin.GET("/popular", h.Popular)
	srv.Gin.POST("/checkout", h.GetCheckoutURL)
	srv.Gin.GET("/receipt/:id", h.Receipt)
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
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Service-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
