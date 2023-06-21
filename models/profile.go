package models

import (
	"time"
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Profile struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"ProfileID"`
	FirstName string    `gorm:"size:50;not null" json:"FirstName"`
	LastName  string    `gorm:"size:50;not null" json:"LastName"`
	UserID    uuid.UUID `gorm:"foreignKey:ID" json:"-"`
	User      User      `gorm:"belongs_to:user;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"User"`
	Addresses []Address `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"Addresses"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

type ProfileInput struct {
	FirstName string    `json:"FirstName" binding:"required"`
	LastName  string    `json:"LastName" binding:"required"`
	UserID    uuid.UUID `json:"UserID"`
}

func (profile *Profile) BeforeSave(scope *gorm.DB) error {
	profile.ID = uuid.NewV4()
	return nil
}

func (profile *Profile) Save() (*Profile, error) {
	if err := DB.Create(&profile).Error; err != nil {
		return &Profile{}, err
	}
	return profile, nil
}

func (profile *Profile) FindProfiles() []Profile {
	var profiles []Profile
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
