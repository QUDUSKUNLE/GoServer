package services


import (
	domain "server/internal/core/domain"
)


func (externalServiceHandler *ServicesHandler) SaveAddress(address domain.Address) error {
	return externalServiceHandler.External.SaveAddress(address)
}

func (externalServiceHandler *ServicesHandler) ReadAddress(AddressID string) (*domain.Address, error) {
	return externalServiceHandler.External.ReadAddress(AddressID)
}

func (externalServiceHandler *ServicesHandler) ReadAddresses() ([]*domain.Address, error) {
	return externalServiceHandler.External.ReadAddresses()
}
