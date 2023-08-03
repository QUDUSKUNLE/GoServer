package models

import (
	"errors"
	"html"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"userID"`
	Email     string    `gorm:"size:255;unique" json:"email"`
	Password  string    `gorm:"size:255;" json:"-"`
	Role      string    `gorm:"type:string;default:customer" json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserInputDto struct {
	Email    string `json:"Email" binding:"required,email,lte=100"`
	Password string `json:"Password" binding:"required,gte=6,lte=20"`
}

func (user *User) BeforeSave(tx *gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))
	user.ID = uuid.NewV4()
	return nil
}

func (user *User) Save() error {
	if err := DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (user *User) ValidatePassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return errors.New("incorrect log in credentials")
	}
	return nil
}

func (user *User) FindUserByEmail(email string) (*User, error) {
	if err := DB.Where(&User{Email: email}).First(&user).Error; err != nil {
		return user, errors.New("incorrect log in credentials")
	}
	return user, nil
}

func FindUserByID(ID string) (User, error) {
	user := User{}
	if err := DB.Where("ID = ?", ID).First(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}
