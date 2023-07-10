package models

import (
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"productID"`
	Quantity  int       `gorm:"not null" json:"productQuantity"`
	StockID   uuid.UUID `gorm:"foreignKey:ID"`
	Stock     Stock     `gorm:"belongs_to:stock" json:"stock"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ProductInput struct {
	Quantity int    `json:"quantity" binding:"required"`
	StockID  string `json:"stockID" binding:"required"`
}

func (product *Product) BeforeSave(scope *gorm.DB) error {
	product.ID = uuid.NewV4()
	return nil
}
