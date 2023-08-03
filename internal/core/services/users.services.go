package services

import (
	domain "server/internal/core/domain"
)

func (servicesHandler *ServicesHandler) SaveUser(user domain.User) error {
	return servicesHandler.repository.SaveUser(user)
}

func (servicesHandler *ServicesHandler) ReadUser(UserID string) (*domain.User, error) {
	return servicesHandler.repository.ReadUser(UserID)
}
