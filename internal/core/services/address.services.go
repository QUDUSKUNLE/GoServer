package services

import (
	domain "server/internal/core/domain"
)

func (externalServiceHandler *ServicesHandler) SaveAddress(address domain.Address) error {
	return externalServiceHandler.Internal.SaveAddress(address)
}

func (externalServiceHandler *ServicesHandler) ReadAddress(AddressID string) (*domain.Address, error) {
	return externalServiceHandler.Internal.ReadAddress(AddressID)
}

func (externalServiceHandler *ServicesHandler) ReadAddresses() ([]*domain.Address, error) {
	return externalServiceHandler.Internal.ReadAddresses()
}
