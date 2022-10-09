package server

import "campfirereads/internal/config"

type App struct {
	Config *config.Config
}

func New(configPrefix string) (*App, error) {
	cfg, err := config.Load(configPrefix)
	if err != nil {
		return nil, err
	}

	return &App{
		Config: cfg,
	}, nil
}
