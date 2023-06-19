package models

import (
	"time"
	"gorm.io/gorm"
	"github.com/satori/go.uuid"
)

type Stock struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"StockID"`
	Type string `gorm:"size:10;not null;" json:"Type"`
	Cost float32 `gorm:"not null" json:"-"`
	Price float32 `gorm:"not null" json:"Price"`
	Province string `gorm:"not null;unique" json:"Province"`
	Unit int `gorm:"not null" json:"-"`
	Slot int `gorm:"not null" json:"Slot"`
	Description string `gorm:"size:255;not null" json:"Description"`
	CreatedAt time.Time `json:"CreatedAt"`
  UpdatedAt time.Time `json:"UpdatedAt"`
}

type CreateStockInput struct {
	Type string `json:"Type" binding:"required"`
	Description string `json:"Description" binding:"required"`
	Province string `json:"Province" binding:"required"`
	Cost float32 `json:"Cost" binding:"required"`
	Price float32 `json:"Price"`
	Unit int `json:"Unit" binding:"required"`
	Slot int `json:"Slot"`
}

func (stock *Stock) BeforeSave(scope *gorm.DB) error {
	stock.ID = uuid.NewV4()
	return nil
}

 func (stock *Stock) Save() (*Stock, error) {
	if err := DB.Create(&stock).Error; err != nil {
		return &Stock{}, err
	}
	return stock, nil
}

func (stock *Stock) FindAll() []Stock {
	var stocks []Stock
	DB.Find(&stocks)
	return stocks
}

func (stock *Stock) FindIn(IDs []string) []Stock {
	var stocks []Stock
	DB.Where(IDs).Find(&stocks)
	return stocks
}

func (stock *Stock) FindStockByID(ID string) (*Stock, error) {
	if err := DB.Where("id = ?", ID).First(&stock).Error; err != nil {
		return stock, err
	}
	return stock, nil
}

func (stock *Stock) FindStockBy(ID string) (Stock, error) {
	if err := DB.Where("id = ?", ID).First(&stock).Error; err != nil {
		return *stock, err
	}
	return *stock, nil
}
