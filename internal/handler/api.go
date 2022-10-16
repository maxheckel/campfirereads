package handler

import (
	"github.com/maxheckel/campfirereads/internal/service"
)

type APIHandler struct {
	google service.Google
	amazon service.Amazon
}

func NewAPI(g service.Google, a service.Amazon) APIHandler {
	return APIHandler{google: g, amazon: a}
}
