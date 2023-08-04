package domain

import (
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Profile struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"profileID"`
	FirstName string    `gorm:"size:50;not null" json:"firstName"`
	LastName  string    `gorm:"size:50;not null" json:"lastName"`
	UserID    uuid.UUID `gorm:"foreignKey:ID" json:"-"`
	User      User      `gorm:"belongs_to:user;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	Addresses []Address `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"addresses"`
}

type ProfileInputDto struct {
	FirstName string    `json:"firstName" binding:"required"`
	LastName  string    `json:"lastName" binding:"required"`
	UserID    uuid.UUID `json:"userID"`
}

func (profile *Profile) BeforeSave(scope *gorm.DB) error {
	profile.ID = uuid.NewV4()
	return nil
}
