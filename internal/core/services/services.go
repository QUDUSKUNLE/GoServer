package services

import (
	repository "server/internal/adapters/repository"
)

type ServicesHandler struct {
	repository repository.PostgresRepository
}

func ServicesAdapter(service repository.PostgresRepository) *ServicesHandler {
	return &ServicesHandler{
		repository: service,
	}
}
