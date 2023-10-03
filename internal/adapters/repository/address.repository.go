package repository

import (
	domain "server/internal/core/domain"
)

func (repo *PostgresRepository) SaveAddress(address domain.Address) error {
	if err := repo.db.Create(&address).Error; err != nil {
		return err
	}
	return nil
}

func (repo *PostgresRepository) ReadAddress(AddressID string) (*domain.Address, error) {
	address := &domain.Address{}
	if err := repo.db.First(&address, "id = ?", AddressID).Error; err != nil {
		return &domain.Address{}, err
	}
	return address, nil
}

func (repo *PostgresRepository) ReadAddresses() ([]*domain.Address, error) {
	var addresses []*domain.Address
	repo.db.Find(&addresses)
	return addresses, nil
}
