package domain

import (
	"time"
	"github.com/satori/go.uuid"
)

type Address struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key" json:"AddressID"`
	StreetNo   int       `gorm:"not null" sql:"unique:idx_streetno_streetname" json:"StreetNo"`
	StreetName string    `gorm:"not null" sql:"unique:idx_streetno_streetname" json:"StreetName"`
	Province   string    `gorm:"not null" json:"Province"`
	State      string    `gorm:"not null" json:"State"`
	ProfileID  uuid.UUID
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

type AddressDto struct {
	StreetNo   int       `json:"streetNo" binding:"required,gte=0,lte=1000"`
	StreetName string    `json:"streetName" binding:"required,max=50"`
	Province   string    `json:"province" binding:"required,max=50"`
	State      string    `json:"state" binding:"required,max=50"`
	UserID     uuid.UUID `json:"userID"`
}
