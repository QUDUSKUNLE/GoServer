package handlers

import (
	services "server/internal/core/services"
)

type HTTPHandler struct {
	ExternalServicesAdapter  services.ServicesHandler
	InternalServicesAdapter  services.ServicesHandler
}

func HTTPAdapter(services services.ServicesHandler) *HTTPHandler {
	return &HTTPHandler{
		ExternalServicesAdapter: services,
		InternalServicesAdapter: services,
	}
}
