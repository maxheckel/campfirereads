package server

import (
	"campfirereads/internal/handler"
	"campfirereads/internal/service"
)

func NewAPI() (*App, error) {
	srv, err := New("API")
	if err != nil {
		panic(err)
	}
	h := handler.NewAPI(service.NewGoogle(srv.Config))
	srv.Gin.GET("/search", h.Search)

	return srv, nil
}
