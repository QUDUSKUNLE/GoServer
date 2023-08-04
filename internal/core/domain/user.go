package domain

import (
	"errors"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"html"
	"strings"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"UserID"`
	Email     string    `gorm:"size:255;unique" json:"Email"`
	Password  string    `gorm:"size:255;" json:"-"`
	Role      string    `gorm:"type:string;default:customer" json:"Role"`
}

type UserInputDto struct {
	Email    string `json:"Email" binding:"required,email,lte=100"`
	Password string `json:"Password" binding:"required,gte=6,lte=20"`
}

func (user *User) BeforeSave(tx *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))
	user.ID = uuid.NewV4()
	return nil
}

func (user *User) ValidatePassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return errors.New("incorrect log in credentials")
	}
	return nil
}
