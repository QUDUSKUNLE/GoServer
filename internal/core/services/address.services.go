package services


import (
	domain "server/internal/core/domain"
)


func (servicesHandler *ServicesHandler) SaveAddress(address *domain.Address) error {
	return servicesHandler.repository.SaveAddress(address)
}

func (servicesHandler *ServicesHandler) ReadAddress(AddressID string) (*domain.Address, error) {
	return servicesHandler.repository.ReadAddress(AddressID)
}

func (servicesHandler *ServicesHandler) ReadAddresses() ([]*domain.Address, error) {
	return servicesHandler.repository.ReadAddresses()
}
