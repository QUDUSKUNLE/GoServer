package repository

import (
	"fmt"
	"errors"
	domain "server/internal/core/domain"
)

func (repo *PostgresRepository) SaveUser(user domain.User) error {
	req := repo.db.Create(&user)
	if req.RowsAffected == 0 {
		return errors.New(fmt.Sprintf("messages not found: %v", req.Error))
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
