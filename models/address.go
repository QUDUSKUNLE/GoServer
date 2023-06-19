package models

import (
	"time"
	"gorm.io/gorm"
	"github.com/satori/go.uuid"
)

type Address struct {
	ID 						uuid.UUID 	`gorm:"type:uuid;primary_key" json:"AddressID"`
	StreetNo 			int 				`gorm:"not null" json:"StreetNo"`
	StreetName 		string  		`gorm:"not null" json:"StreetName"`
	Province 			string  		`gorm:"not null" json:"Province"`
	State 				string  		`gorm:"not null" json:"State"`
	ProfileID 		uuid.UUID
	CreatedAt 		time.Time 	`json:"CreatedAt"`
  UpdatedAt 		time.Time 	`json:"UpdatedAt"`
}

type AddressInput struct {
	StreetNo 			int 				`json:"StreetNo" binding:"required,gte=0,lte=1000"`
	StreetName 		string  		`json:"StreetName" binding:"required,max=50"`
	Province 			string  		`json:"Province" binding:"required,max=50"`
	State 				string  		`json:"State" binding:"required,max=50"`
	UserID 				uuid.UUID   `json:"UserID"`
}

func (address *Address) BeforeSave(scope *gorm.DB) error {
	address.ID = uuid.NewV4()
	return nil
}

func (address *Address) Save() (*Address, error) {
	if err := DB.Create(&address).Error; err != nil {
		return &Address{}, err
	}
	return address, nil
}

func (address *Address) FindAll() []Address {
	var addresses []Address
	DB.Find(&addresses)
	return addresses
}
