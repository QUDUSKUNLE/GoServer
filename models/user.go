package models

import (
	"errors"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"html"
	"strings"
	"time"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"UserID"`
	Username  string    `gorm:"size:255;unique" json:"Username"`
	Password  string    `gorm:"size:255;" json:"-"`
	Role      string    `gorm:"type:string;default:customer" json:"Role"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

type UserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (user *User) BeforeSave(tx *gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	user.ID = uuid.NewV4()
	return nil
}

func (user *User) Save() (*User, error) {
	if err := DB.Create(&user).Error; err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) ValidatePassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return errors.New("incorrect log in credentials")
	}
	return nil
}

func (user *User) FindUserByUsername(username string) (*User, error) {
	if err := DB.Where(&User{Username: username}).First(&user).Error; err != nil {
		return user, errors.New("incorrect log in credentials")
	}
	return user, nil
}

func FindUserByID(ID string) (User, error) {
	var user User
	if err := DB.Where("ID = ?", ID).First(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}
