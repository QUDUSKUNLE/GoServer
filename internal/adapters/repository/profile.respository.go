package repository

import (
	"fmt"
	"errors"
	"gorm.io/gorm/clause"
	domain "server/internal/core/domain"
)

func (repo *PostgresRepository) SaveProfile(profile domain.Profile) error {
	req := repo.db.Create(&profile)
	if req.RowsAffected == 0 {
		return errors.New(fmt.Sprintf("error creating profile: %v", req.Error))
	}
	return nil
}

func (repo *PostgresRepository) ReadProfile(ProfileID string) (*domain.Profile, error) {
	profile := &domain.Profile{}
	req := repo.db.First(&profile, "id = ?", ProfileID)
	if req.RowsAffected == 0 {
		return nil, errors.New("profile not found")
	}
	return profile, nil
}

func (repo *PostgresRepository) ReadProfiles() ([]*domain.Profile, error) {
	var profiles []*domain.Profile
	req := repo.db.Find(&profiles)
	if req.Error != nil {
		return nil, errors.New("no profile found")
	}
	return profiles, nil
}

func (repo *PostgresRepository) ReadProfileByUserID(UserID string) (*domain.Profile, error) {
	var profile *domain.Profile
	req := repo.db.Preload(clause.Associations).Where("user_id = ?", UserID).First(&profile)
	if req.RowsAffected == 0 {
		return nil, errors.New("no profile found")
	}
	return profile, nil
}
