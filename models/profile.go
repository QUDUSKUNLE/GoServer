package models

import (
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type Profile struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"profileID"`
	FirstName string    `gorm:"size:50;not null" json:"firstName"`
	LastName  string    `gorm:"size:50;not null" json:"lastName"`
	UserID    uuid.UUID `gorm:"foreignKey:ID" json:"-"`
	User      User      `gorm:"belongs_to:user;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	Addresses []Address `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"addresses"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ProfileInput struct {
	FirstName string    `json:"firstName" binding:"required"`
	LastName  string    `json:"lastName" binding:"required"`
	UserID    uuid.UUID `json:"userID"`
}

func (profile *Profile) BeforeSave(scope *gorm.DB) error {
	profile.ID = uuid.NewV4()
	return nil
}

func (profile *Profile) Save() error {
	if err := DB.Create(&profile).Error; err != nil {
		return err
	}
	return nil
}

func (profile *Profile) FindProfiles() []Profile {
	profiles := []Profile{}
	DB.Preload(clause.Associations).Find(&profiles)
	return profiles
}

func (profile *Profile) FindProfile(ID string) (*Profile, error) {
	if err := DB.Preload(clause.Associations).Where("id = ?", ID).First(&profile).Error; err != nil {
		return &Profile{}, err
	}
	return profile, nil
}

func (profile *Profile) FindProfileByUserID(UserID string) (*Profile, error) {
	if err := DB.Preload(clause.Associations).Where("user_id = ?", UserID).First(&profile).Error; err != nil {
		return &Profile{}, err
	}
	return profile, nil
}
