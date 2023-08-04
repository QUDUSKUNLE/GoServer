package repository

import (
	"fmt"
	"errors"
	domain "server/internal/core/domain"
)

func (repo *PostgresRepository) SaveAddress(address domain.Address) error {
	req := repo.db.Create(&address)
	if req.RowsAffected == 0 {
		return errors.New(fmt.Sprintf("error creating address: %v", req.Error))
	}
	return nil
}

func (repo *PostgresRepository) ReadAddress(AddressID string) (*domain.Address, error) {
	address := &domain.Address{}
	req := repo.db.First(&address, "id = ?", AddressID)
	if req.RowsAffected == 0 {
		return nil, errors.New("address not found")
	}
	return address, nil
}

func (repo *PostgresRepository) ReadAddresses() ([]*domain.Address, error) {
	var addresses []*domain.Address
	req := repo.db.Find(&addresses)
	if req.Error != nil {
		return nil, errors.New("no address found")
	}
	return addresses, nil
}
