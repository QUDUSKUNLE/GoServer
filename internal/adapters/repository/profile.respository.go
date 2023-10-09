package repository

import (
	domain "server/internal/core/domain"
)

func (repo *PostgresRepository) SaveProfile(profile domain.Profile) error {
	_, err := repo.db.NewInsert().Model(&profile).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (repo *PostgresRepository) ReadProfile(ProfileID string) (*domain.Profile, error) {
	profile := new(domain.Profile)
	if err := repo.db.NewSelect().Model(profile).Where("ID = ?", ProfileID).Scan(ctx); err != nil {
		return &domain.Profile{}, err
	}
	return profile, nil
}

func (repo *PostgresRepository) ReadProfiles() ([]*domain.Profile, error) {
	var profiles []*domain.Profile
	repo.db.NewSelect().Model(&profiles).Limit(20).Scan(ctx)
	return profiles, nil
}

func (repo *PostgresRepository) ReadProfileByUserID(UserID string) (*domain.Profile, error) {
	profile := new(domain.Profile)
	if err := repo.db.NewSelect().Model(profile).Where("user_id = ?", UserID).Scan(ctx); err != nil {
		return &domain.Profile{}, err
	}
	return profile, nil
}
