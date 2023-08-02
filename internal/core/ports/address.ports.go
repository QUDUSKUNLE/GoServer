package ports

import (
	"server/internal/core/domain"
)

type AddressRepository interface {
	SaveAddress(address domain.Address) error
	ReadAddress(AddressID string) (*domain.Address, error)
	ReadAddresses() ([]*domain.Address, error)
	DeleteAddress(AddressID string) (bool, error)
}
