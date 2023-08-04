package services


import (
	domain "server/internal/core/domain"
)


func (servicesHandler *ServicesHandler) SaveAddress(address domain.Address) error {
	return servicesHandler.repo.SaveAddress(address)
}

func (servicesHandler *ServicesHandler) ReadAddress(AddressID string) (*domain.Address, error) {
	return servicesHandler.repo.ReadAddress(AddressID)
}

func (servicesHandler *ServicesHandler) ReadAddresses() ([]*domain.Address, error) {
	return servicesHandler.repo.ReadAddresses()
}
