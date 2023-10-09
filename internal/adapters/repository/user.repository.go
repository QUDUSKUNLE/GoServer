package repository

import (
	domain "server/internal/core/domain"
)

func (repo *PostgresRepository) SaveUser(user domain.User) error {
	_, err := repo.db.NewInsert().Model(&user).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (repo *PostgresRepository) ReadUser(UserID string) (*domain.User, error) {
	user := new(domain.User)
	if err := repo.db.NewSelect().Model(user).Where("ID = ?", UserID).Scan(ctx); err != nil {
		return &domain.User{}, err
	}
	return user, nil
}

func (repo *PostgresRepository) ReadUsers() ([]*domain.User, error) {
	var users []*domain.User
	repo.db.NewSelect().Model(users).Limit(20)
	return users, nil
}

func (repo *PostgresRepository) ReadUserByEmail(Email string) (*domain.User, error) {
	user := new(domain.User)
	if err := repo.db.NewSelect().Model(user).Where("Email = ? ", Email).Scan(ctx); err != nil {
		return &domain.User{}, err
	}
	return user, nil
}
