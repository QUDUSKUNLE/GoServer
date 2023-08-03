package handlers

import (
	services "server/internal/core/services"
)

type HTTPHandler struct {
	svc  services.ServicesHandler
}

func NewHTTPHandlers(services services.ServicesHandler) *HTTPHandler {
	return &HTTPHandler{
		svc: services,
	}
}
