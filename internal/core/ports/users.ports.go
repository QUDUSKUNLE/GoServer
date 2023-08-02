package ports

import (
	"server/internal/core/domain"
)

type UserRepository interface {
	SaveUser(user domain.User) error
	ReadUser(UserID string) (*domain.User, error)
	ReadUsers() ([]*domain.User, error)
	DeleteUser(UserID string) (bool, error)
}
