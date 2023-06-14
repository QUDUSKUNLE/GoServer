package models

import (
	"errors"
	"html"
	"strings"
	"time"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"UserID"`
	Username  string    `gorm:"size:255;not null;unique" json:"Username"`
	Password  string    `gorm:"size:255;not null;" json:"-"`
	Orders 		[]Order  	`gorm:"foreignKey:UserID" json:"-"`
 	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type UserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (user *User) Save() (*User, error) {
	if err := DB.Create(&user).Error; err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	user.ID = uuid.NewV4()
	return nil
}

func (user *User) ValidatePassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return errors.New("invalid log in credentials")
	}
	return nil
}

func (user *User) FindUserByUsername(username string) (*User, error) {
	if err := DB.Where(&User{Username: username}).First(&user).Error; err != nil {
		return user, errors.New("invalid log in credentials")
	}
	return user, nil
}

func FindUserById(ID string) (User, error) {
	var user User
	if err := DB.Where("ID = ?", ID).First(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}
