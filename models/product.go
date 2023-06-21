package models

import (
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"ProductID"`
	Quantity  int       `gorm:"not null" json:"ProductQuantity"`
	StockID   uuid.UUID `gorm:"foreignKey:ID"`
	Stock     Stock     `gorm:"belongs_to:stock" json:"Stock"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

type ProductInput struct {
	Quantity int   `json:"Quantity" binding:"required"`
	StockID  string `json:"StockID" binding:"required"`
}

func (product *Product) BeforeSave(scope *gorm.DB) error {
	product.ID = uuid.NewV4()
	return nil
}
