package repository

import (
	"gorm.io/gorm/clause"
	domain "server/internal/core/domain"
)

func (repo *PostgresRepository) SaveProfile(profile domain.Profile) error {
	if err := repo.db.Create(&profile).Error; err != nil {
		return err
	}
	return nil
}

func (repo *PostgresRepository) ReadProfile(ProfileID string) (*domain.Profile, error) {
	profile := &domain.Profile{}
	if err := repo.db.Preload(clause.Associations).Where("id = ?", ProfileID).First(&profile).Error; err != nil {
		return &domain.Profile{}, err
	}
	return profile, nil
}

func (repo *PostgresRepository) ReadProfiles() ([]*domain.Profile, error) {
	var profiles []*domain.Profile
	repo.db.Preload(clause.Associations).Find(&profiles)
	return profiles, nil
}

func (repo *PostgresRepository) ReadProfileByUserID(UserID string) (*domain.Profile, error) {
	var profile *domain.Profile
	if err := repo.db.Preload(clause.Associations).Where("user_id = ?", UserID).First(&profile).Error; err != nil {
		return &domain.Profile{}, err
	}
	return profile, nil
}
