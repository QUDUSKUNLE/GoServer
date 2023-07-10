package models

import (
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type Address struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key" json:"addressID"`
	StreetNo   int       `gorm:"not null" sql:"unique:idx_streetno_streetname" json:"streetNo"`
	StreetName string    `gorm:"not null" sql:"unique:idx_streetno_streetname" json:"streetName"`
	Province   string    `gorm:"not null" json:"province"`
	State      string    `gorm:"not null" json:"state"`
	ProfileID  uuid.UUID
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type AddressInput struct {
	StreetNo   int       `json:"streetNo" binding:"required,gte=0,lte=1000"`
	StreetName string    `json:"streetName" binding:"required,max=50"`
	Province   string    `json:"province" binding:"required,max=50"`
	State      string    `json:"state" binding:"required,max=50"`
	UserID     uuid.UUID `json:"userID"`
}

func (address *Address) BeforeSave(scope *gorm.DB) error {
	address.ID = uuid.NewV4()
	return nil
}

func (address *Address) Save() error {
	if err := DB.Create(&address).Error; err != nil {
		return err
	}
	return nil
}

func (address *Address) FindAddresses() []Address {
	var addresses []Address
	DB.Preload(clause.Associations).Find(&addresses)
	return addresses
}

func (address *Address) FindAddress(ID string) (*Address, error) {
	if err := DB.Preload(clause.Associations).Where("id = ?", ID).First(&address).Error; err != nil {
		return &Address{}, err
	}
	return address, nil
}
