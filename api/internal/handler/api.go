package handler

import (
	"github.com/maxheckel/campfirereads/internal/service"
	"github.com/maxheckel/campfirereads/internal/service/cache"
	"github.com/maxheckel/campfirereads/internal/service/payments"
)

type APIHandler struct {
	google   service.Google
	amazon   service.Amazon
	nyt      service.NYT
	payments payments.Service
	cache    cache.Service
}

func NewAPI(g service.Google, a service.Amazon, nyt service.NYT, cache cache.Service, payments payments.Service) APIHandler {
	return APIHandler{google: g, amazon: a, nyt: nyt, cache: cache, payments: payments}
}
