package models

import (
	"time"

	"gorm.io/gorm"
	"github.com/satori/go.uuid"
)

type Profile struct {
	ID        	uuid.UUID 	`gorm:"type:uuid;primary_key" json:"UserID"`
	Email  			string    	`gorm:"size:50;unique" json:"Email"`
	FirstName  	string    	`gorm:"size:50;not null" json:"FirstName"`
	LastName 		string      `gorm:"size:50;not null" json:"LastName"`
	UserID 		  uuid.UUID
	Addresses 	[]Address   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"Addresses"`
 	CreatedAt 	time.Time 	`json:"CreatedAt"`
	UpdatedAt 	time.Time 	`json:"UpdatedAt"`
}

type ProfileInput struct {
	Email  			string    	`json:"Email" binding:"required"`
	FirstName  	string    	`json:"FirstName" binding:"required"`
	LastName 		string      `json:"LastName" binding:"required"`
	UserID 		  uuid.UUID   `json:"UserID"`
}

func (profile *Profile) BeforeSave(scope *gorm.DB) error {
	profile.ID = uuid.NewV4()
	return nil
}

func (user *Profile) Save() (*Profile, error) {
	if err := DB.Create(&user).Error; err != nil {
		return &Profile{}, err
	}
	return user, nil
}
