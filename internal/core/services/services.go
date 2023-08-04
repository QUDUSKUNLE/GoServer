package services

import (
	ports "server/internal/core/ports"
)

type ServicesHandler struct {
	repo ports.ServerRepository
}

func ServicesAdapter(repo ports.ServerRepository) *ServicesHandler {
	return &ServicesHandler{
		repo: repo,
	}
}
