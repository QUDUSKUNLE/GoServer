package services

import (
	ports "server/internal/core/ports"
	domain "server/internal/core/domain"
)

type UserService struct {
	userRepository ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{
		userRepository: repo,
	}
}

func (u *UserService) SaveUser(user domain.User) error {
	return u.userRepository.SaveUser(user)
}

func (u *UserService) ReadUser(UserID string) (*domain.User, error) {
	return u.userRepository.ReadUser(UserID)
}

func (u *UserService) ReadUsers() ([]*domain.User, error) {
	return u.userRepository.ReadUsers()
}

func (u *UserService) DeleteUser(UserID string) (bool, error) {
	return u.userRepository.DeleteUser(UserID)
}
