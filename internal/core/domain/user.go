package domain

import (
	"time"
	"errors"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uuid.UUID 	`gorm:"type:uuid;primaryKey" json:"UserID"`
	Email     string    	`gorm:"size:255;unique" json:"Email"`
	Password  string    	`gorm:"size:255;" json:"-"`
	Role      string    	`gorm:"type:string;default:customer" json:"Role"`
	Profile   Profile
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserDto struct {
	Email    string `json:"Email" binding:"required,email,lte=100"`
	Password string `json:"Password" binding:"required,gte=6,lte=20"`
}

func (user *User) ValidatePassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return errors.New("incorrect log in credentials")
	}
	return nil
}
