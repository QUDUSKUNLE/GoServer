package domain

import (
	"github.com/satori/go.uuid"
	"time"
)

type Profile struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"ProfileID"`
	FirstName string    `gorm:"size:50;index:idx_full_name" json:"FirstName"`
	LastName  string    `gorm:"size:50;index:idx_full_name" json:"LastName"`
	UserID    *uuid.UUID `gorm:"foreignKey:ID" json:"-"`
	User      *User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"User"`
	Addresses []Address `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"Addresses"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

type ProfileDto struct {
	FirstName string    `json:"FirstName" binding:"required"`
	LastName  string    `json:"LastName" binding:"required"`
	UserID    uuid.UUID `json:"UserID"`
}
