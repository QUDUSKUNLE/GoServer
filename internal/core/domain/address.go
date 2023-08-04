package domain

import (
	"github.com/satori/go.uuid"
)

type Address struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key" json:"addressID"`
	StreetNo   int       `gorm:"not null" sql:"unique:idx_streetno_streetname" json:"streetNo"`
	StreetName string    `gorm:"not null" sql:"unique:idx_streetno_streetname" json:"streetName"`
	Province   string    `gorm:"not null" json:"province"`
	State      string    `gorm:"not null" json:"state"`
	ProfileID  uuid.UUID
}
