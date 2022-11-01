package handler

import (
	"github.com/maxheckel/campfirereads/internal/service"
	"github.com/maxheckel/campfirereads/internal/service/cache"
)

type APIHandler struct {
	google service.Google
	amazon service.Amazon
	nyt    service.NYT
	cache  cache.Cache
}

func NewAPI(g service.Google, a service.Amazon, nyt service.NYT, cache cache.Cache) APIHandler {
	return APIHandler{google: g, amazon: a, nyt: nyt, cache: cache}
}
