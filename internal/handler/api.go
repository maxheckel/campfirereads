package handler

import "campfirereads/internal/service"

type APIHandler struct {
	google service.Google
}

func NewAPI(g service.Google) APIHandler {
	return APIHandler{google: g}
}
