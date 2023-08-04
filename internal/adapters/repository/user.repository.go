package repository

import (
	"fmt"
	"errors"
	domain "server/internal/core/domain"
)

func (repo *PostgresRepository) SaveUser(user domain.User) error {
	req := repo.db.Create(&user)
	if req.RowsAffected == 0 {
		return errors.New(fmt.Sprintf("%v", req.Error))
	}
	return nil
}

func (repo *PostgresRepository) ReadUser(UserID string) (*domain.User, error) {
	user := &domain.User{}
	req := repo.db.First(&user, "id = ?", UserID)
	if req.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (repo *PostgresRepository) ReadUsers() ([]*domain.User, error) {
	var users []*domain.User
	req := repo.db.Find(&users)
	if req.Error != nil {
		return nil, errors.New("no user found")
	}
	return users, nil
}

func (repo *PostgresRepository) ReadUserByEmail(Email string) (*domain.User, error) {
	user := &domain.User{}
	req := repo.db.Where(&domain.User{Email: Email}).First(&user)
	if req.RowsAffected == 0 {
		return nil, errors.New("no user found")
	}
	return user, nil
}
