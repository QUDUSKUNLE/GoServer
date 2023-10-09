package domain

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type Profile struct {
	bun.BaseModel 		   `bun:"table:profiles"`

	ID        uuid.UUID 	`bun:",pk,type:uuid,default:uuid_generate_v4()" json:"ProfileID"`
	FirstName string    	`json:"FirstName"`
	LastName  string    	`json:"LastName"`
	UserID    uuid.UUID   `bun:"type:uuid"`
	User      User        `bun:"rel:belongs-to,join:user_id=id"`
	Addresses []Address 	`json:"Addresses"`
	CreatedAt time.Time   `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time   `bun:",nullzero,notnull,default:current_timestamp"`
}

type ProfileDto struct {
	FirstName string    `json:"FirstName" binding:"required"`
	LastName  string    `json:"LastName" binding:"required"`
	UserID    uuid.UUID `json:"UserID"`
}
