package repository

import (
	domain "server/internal/core/domain"
)

func (repo *PostgresRepository) SaveAddress(address domain.Address) error {
	_, err := repo.db.NewInsert().Model(address).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (repo *PostgresRepository) ReadAddress(AddressID string) (*domain.Address, error) {
	address := new(domain.Address)
	if err := repo.db.NewSelect().Model(address).Where("id = ?", AddressID).Scan(ctx); err != nil {
		return &domain.Address{}, err
	}
	return address, nil
}

func (repo *PostgresRepository) ReadAddresses() ([]*domain.Address, error) {
	var addresses []*domain.Address
	repo.db.NewSelect().Model(addresses).Limit(20)
	return addresses, nil
}
