package handlers

import (
	"server/internal/core/services"
)

type HTTPHandler struct {
	user  services.UserService
}

func NewHTTPHandlers(services HTTPHandler) *HTTPHandler {
	return &HTTPHandler{
		user: services.user,
	}
}
