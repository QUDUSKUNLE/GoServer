package ports

import (
	domain "server/internal/core/domain"
)

type ServerService interface {
	SaveUser(user domain.User) error
	ReadUser(UserID string) (*domain.User, error)
	ReadUsers() ([]*domain.User, error)
	ReadUserByEmail(Email string) (*domain.User, error)

	SaveAddress(address domain.Address) error
	ReadAddress(AddressID string) (*domain.Address, error)
	ReadAddresses() ([]*domain.Address, error)
}

type ServerRepository interface {
	SaveUser(user domain.User) error
	ReadUser(UserID string) (*domain.User, error)
	ReadUsers() ([]*domain.User, error)
	ReadUserByEmail(Email string) (*domain.User, error)

	SaveAddress(address domain.Address) error
	ReadAddress(AddressID string) (*domain.Address, error)
	ReadAddresses() ([]*domain.Address, error)
}
