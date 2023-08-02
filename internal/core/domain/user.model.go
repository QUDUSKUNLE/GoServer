package domain

import (
	"github.com/satori/go.uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"userID"`
	Email     string    `gorm:"size:255;unique" json:"email"`
	Password  string    `gorm:"size:255;" json:"-"`
	Role      string    `gorm:"type:string;default:customer" json:"role"`
}
