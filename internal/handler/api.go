package handler

import (
	"github.com/maxheckel/campfirereads/internal/service"
)

type APIHandler struct {
	google service.Google
	amazon service.Amazon
	nyt    service.NYT
}

func NewAPI(g service.Google, a service.Amazon, nyt service.NYT) APIHandler {
	return APIHandler{google: g, amazon: a, nyt: nyt}
}
