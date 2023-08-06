package handlers

import (
	services "server/internal/core/services"
)

type HTTPHandler struct {
	ServicesAdapter  services.ServicesHandler
}

func HTTPAdapter(services services.ServicesHandler) *HTTPHandler {
	return &HTTPHandler{
		ServicesAdapter: services,
	}
}
