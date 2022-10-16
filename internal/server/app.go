package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/maxheckel/campfirereads/internal/config"
)

type App struct {
	Config *config.Config
	Gin    *gin.Engine
}

func New(configPrefix string) (*App, error) {
	cfg, err := config.Load(configPrefix)
	if err != nil {
		return nil, err
	}

	return &App{
		Config: cfg,
		Gin:    gin.Default(),
	}, nil
}

func (a *App) Start() error {
	return a.Gin.Run(fmt.Sprintf(":%d", a.Config.Port))
}
