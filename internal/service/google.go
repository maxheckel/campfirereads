package service

import "campfirereads/internal/config"

type BookSearch struct {
}

type Google interface {
	GetBooks(search BookSearch)
}

type google struct {
	config *config.Config
}

func NewGoogle(cfg *config.Config) Google {
	return &google{config: cfg}
}

func (g google) GetBooks(search BookSearch) {

}
