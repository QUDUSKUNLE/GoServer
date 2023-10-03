package handlers

import (
	services "server/internal/core/services"
)

type HTTPHandler struct {
	ServicesAdapter  services.ServicesHandler
}

func HTTPAdapter(serviceHandlers services.ServicesHandler) *HTTPHandler {
	return &HTTPHandler{
		ServicesAdapter: serviceHandlers,
	}
}
