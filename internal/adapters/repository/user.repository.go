package repository

import (
	domain "server/internal/core/domain"
)

func (repo *PostgresRepository) SaveUser(user domain.User) error {
	if err := repo.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *PostgresRepository) ReadUser(UserID string) (*domain.User, error) {
	user := &domain.User{}
	if err := repo.db.Where("ID = ?", UserID).First(&user).Error; err != nil {
		return &domain.User{}, err
	}
	return user, nil
}

func (repo *PostgresRepository) ReadUsers() ([]*domain.User, error) {
	var users []*domain.User
	repo.db.Find(&users)
	return users, nil
}

func (repo *PostgresRepository) ReadUserByEmail(Email string) (*domain.User, error) {
	user := &domain.User{}
	if err := repo.db.Where(&domain.User{Email: Email}).First(&user).Error; err != nil {
		return &domain.User{}, err
	}
	return user, nil
}
