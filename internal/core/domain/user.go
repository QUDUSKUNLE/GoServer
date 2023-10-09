package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	bun.BaseModel 		   `bun:"table:users"`

	ID        uuid.UUID 	`bun:",pk,type:uuid,default:uuid_generate_v4()"`
	Email     string    	`bun:"email,unique"`
	Password  string    	`bun:"password"`
	Role      string    	`bun:",nullzero,notnull,default:'customer'"`
	CreatedAt time.Time   `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time   `bun:",nullzero,notnull,default:current_timestamp"`
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
