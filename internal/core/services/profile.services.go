package services

import (
	domain "server/internal/core/domain"
)

func (externalServiceHandler *ServicesHandler) SaveProfile(profile domain.Profile) error {
	return externalServiceHandler.External.SaveProfile(profile)
}

func (externalServiceHandler *ServicesHandler) ReadProfile(ProfileID string) (*domain.Profile, error) {
	return externalServiceHandler.External.ReadProfile(ProfileID)
}

func (externalServiceHandler *ServicesHandler) ReadProfiles() ([]*domain.Profile, error) {
	return externalServiceHandler.External.ReadProfiles()
}

func (externalServiceHandler *ServicesHandler) ReadProfileByUserID(UserID string) (*domain.Profile, error) {
	return externalServiceHandler.External.ReadProfileByUserID(UserID)
}
