package services


import (
	ports "server/internal/core/ports"
	domain "server/internal/core/domain"
)

type AddressService struct {
	addressRepository ports.AddressRepository
}

func NewAddressService(repo ports.AddressRepository) *AddressService {
	return &AddressService{
		addressRepository: repo,
	}
}

func (u *AddressService) SaveAddress(address domain.Address) error {
	return u.addressRepository.SaveAddress(address)
}

func (u *AddressService) ReadAddress(AddressID string) (*domain.Address, error) {
	return u.addressRepository.ReadAddress(AddressID)
}

func (u *AddressService) ReadAddresses() ([]*domain.Address, error) {
	return u.addressRepository.ReadAddresses()
}

func (u *AddressService) DeleteAddress(AddressID string) (bool, error) {
	return u.addressRepository.DeleteAddress(AddressID)
}
