package domain

import (
	"github.com/uptrace/bun"
	"time"
)

type Profile struct {
	bun.BaseModel 		   `bun:"table:profiles"`

	ID        int         `bun:",pk"`
	FirstName string    	`json:"FirstName"`
	LastName  string    	`json:"LastName"`
	UserID    int
	User      *User       `bun:"rel:belongs-to,join:user_id=id"`
	Addresses []Address 	`json:"-"`
	CreatedAt time.Time   `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time   `bun:",nullzero,notnull,default:current_timestamp"`
}

type ProfileDTO struct {
	FirstName string    `json:"FirstName" binding:"required"`
	LastName  string    `json:"LastName" binding:"required"`
	UserID    int       `json:"UserID"`
}
