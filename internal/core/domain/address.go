package domain

import (
	"time"
	"github.com/google/uuid"
)

type Address struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key" json:"AddressID"`
	StreetNo   int       `gorm:"not null" sql:"unique:idx_streetno_streetname" json:"StreetNo"`
	StreetName string    `gorm:"not null" sql:"unique:idx_streetno_streetname" json:"StreetName"`
	Province   string    `gorm:"not null" json:"Province"`
	State      string    `gorm:"not null" json:"State"`
	Profile    Profile
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AddressDto struct {
	StreetNo   int       `json:"StreetNo" binding:"required,gte=0,lte=1000"`
	StreetName string    `json:"StreetName" binding:"required,max=50"`
	Province   string    `json:"Province" binding:"required,max=50"`
	State      string    `json:"State" binding:"required,max=50"`
	UserID     uuid.UUID `json:"UserID"`
}
