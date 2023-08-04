package services

import (
	domain "server/internal/core/domain"
)

func (servicesHandler *ServicesHandler) SaveUser(user domain.User) error {
	return servicesHandler.repo.SaveUser(user)
}

func (servicesHandler *ServicesHandler) ReadUser(UserID string) (*domain.User, error) {
	return servicesHandler.repo.ReadUser(UserID)
}

func (servicesHandler *ServicesHandler) ReadUsers() ([]*domain.User, error) {
	return servicesHandler.repo.ReadUsers()
}

func (servicesHandler *ServicesHandler) ReadUserByEmail(Email string) (*domain.User, error) {
	return servicesHandler.repo.ReadUserByEmail(Email)
}
