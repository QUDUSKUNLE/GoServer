package models

import (
	"time"
	"gorm.io/gorm"
	"github.com/satori/go.uuid"
)

type Product struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key" json:"ProductID"`
	Quantity int `gorm:"not null" json:"ProductQuantity"`
	Stock Stock  `gorm:"" json:"Stock"`
	StockID uuid.UUID `gorm:"foreignKey:ID"`
	CreatedAt time.Time `json:"CreatedAt"`
  UpdatedAt time.Time `json:"UpdatedAt"`
}

type ProductInput struct {
	Quantity int `json:"Quantity" binding:"required"`
	Stock Stock `json:"Stock" binding:"required"`
	StockID uuid.UUID
}


func (product *Product) BeforeSave(scope *gorm.DB) error {
	product.ID = uuid.NewV4()
	return nil
 }

 func (product *Product) Save(productInput []ProductInput) ([]Product, error) {
	var products []Product
	if err := DB.Create(&productInput).Error; err != nil {
		return []Product{}, err
	}
	return products, nil
}
