package ports

import (
	domain "server/internal/core/domain"
)

type ServerServicePorts interface {
	SaveUser(user domain.User) error
	ReadUser(UserID string) (*domain.User, error)
	ReadUsers() ([]*domain.User, error)
	ReadUserByEmail(Email string) (*domain.User, error)

	SaveAddress(address domain.Address) error
	ReadAddress(AddressID string) (*domain.Address, error)
	ReadAddresses() ([]*domain.Address, error)

	SaveProfile(profile domain.Profile) error
	ReadProfile(ProfileID string) (*domain.Profile, error)
	ReadProfiles() ([]*domain.Profile, error)
	ReadProfileByUserID(UserID string) (*domain.Profile, error)
}

type ServerRepositoryPorts interface {
	SaveUser(user domain.User) error
	ReadUser(UserID string) (*domain.User, error)
	ReadUsers() ([]*domain.User, error)
	ReadUserByEmail(Email string) (*domain.User, error)

	SaveAddress(address domain.Address) error
	ReadAddress(AddressID string) (*domain.Address, error)
	ReadAddresses() ([]*domain.Address, error)

	SaveProfile(profile domain.Profile) error
	ReadProfile(ProfileID string) (*domain.Profile, error)
	ReadProfiles() ([]*domain.Profile, error)
	ReadProfileByUserID(UserID string) (*domain.Profile, error)
}

