package services

import (
	domain "server/internal/core/domain"
)

func (externalServiceHandler *ServicesHandler) SaveProfile(profile domain.Profile) error {
	return externalServiceHandler.Internal.SaveProfile(profile)
}

func (externalServiceHandler *ServicesHandler) ReadProfile(ProfileID string) (*domain.Profile, error) {
	return externalServiceHandler.Internal.ReadProfile(ProfileID)
}

func (externalServiceHandler *ServicesHandler) ReadProfiles() ([]*domain.Profile, error) {
	return externalServiceHandler.Internal.ReadProfiles()
}

func (externalServiceHandler *ServicesHandler) ReadProfileByUserID(UserID string) (*domain.Profile, error) {
	return externalServiceHandler.Internal.ReadProfileByUserID(UserID)
}
