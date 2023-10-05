package domain

import (
	"github.com/google/uuid"
	"time"
)

type Profile struct {
	ID        uuid.UUID 	`bun:"type:uuid,pk" json:"ProfileID"`
	FirstName string    	`json:"FirstName"`
	LastName  string    	`json:"LastName"`
	UserID    uuid.UUID   `bun:"type:uuid"`
	User      User        `bun:"rel:belongs-to,join:user_id=id"`
	Addresses []Address 	`json:"Addresses"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ProfileDto struct {
	FirstName string    `json:"FirstName" binding:"required"`
	LastName  string    `json:"LastName" binding:"required"`
	UserID    uuid.UUID `json:"UserID"`
}
