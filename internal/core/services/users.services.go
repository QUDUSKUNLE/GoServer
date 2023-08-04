package services

import (
	domain "server/internal/core/domain"
)

func (externalServiceHandler *ServicesHandler) SaveUser(user domain.User) error {
	return externalServiceHandler.External.SaveUser(user)
}

func (externalServiceHandler *ServicesHandler) ReadUser(UserID string) (*domain.User, error) {
	return externalServiceHandler.External.ReadUser(UserID)
}

func (externalServiceHandler *ServicesHandler) ReadUsers() ([]*domain.User, error) {
	return externalServiceHandler.External.ReadUsers()
}

func (internalServiceHandler *ServicesHandler) ReadUserByEmail(Email string) (*domain.User, error) {
	return internalServiceHandler.Internal.ReadUserByEmail(Email)
}
