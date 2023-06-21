package models

import (
	"errors"
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type Stock struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" json:"StockID"`
	Type         string    `gorm:"size:10;index:idx_type_province,unique" json:"Type"`
	Description  string    `gorm:"size:255;not null" json:"Description"`
	Availability bool      `gorm:"type:bool;default:true" json:"Availability"`
	Cost         float32   `gorm:"not null" json:"-"`
	Price        float32   `gorm:"not null" json:"Price"`
	Province     string    `gorm:"size:15;index:idx_type_province,unique" json:"Province"`
	Unit         int       `gorm:"not null" json:"-"`
	Slot         int       `gorm:"not null" json:"Slot"`
	CreatedAt    time.Time `json:"CreatedAt"`
	UpdatedAt    time.Time `json:"UpdatedAt"`
}

type CreateStockInput struct {
	Type        string  `json:"Type" binding:"required"`
	Description string  `json:"Description" binding:"required"`
	Province    string  `json:"Province" binding:"required"`
	Cost        float32 `json:"Cost" binding:"required"`
	Price       float32 `json:"Price"`
	Unit        int     `json:"Unit" binding:"required"`
	Slot        int     `json:"Slot"`
}

type UpdateStockInput struct {
	Type        string  `json:"Type"`
	Description string  `json:"Description"`
	Province    string  `json:"Province"`
	Cost        float32 `json:"Cost"`
	Price       float32 `json:"Price"`
	Unit        int     `json:"Unit"`
	Slot        int     `json:"Slot"`
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
		return &Stock{}, err
	}
	return stock, nil
}

func (stock *Stock) FindStockBy(ID string) (Stock, error) {
	if err := DB.Where("id = ?", ID).First(&stock).Error; err != nil {
		return *stock, err
	}
	return *stock, nil
}

func (stock *Stock) Update(updateStock UpdateStockInput, id string) (*Stock, error) {
	if err := DB.First(&stock, id).Error; err != nil {
		return &Stock{}, err
	}
	if err := DB.Model(&stock).Updates(updateStock).Error; err != nil {
		return &Stock{}, err
	}
	return stock, nil
}

func (stock *Stock) Delete(ID string) (bool, error) {
	if err := DB.Where("id = ?", ID).First(&stock).Error; err != nil {
		return false, errors.New("record not found")
	}
	DB.Delete(&stock)
	return true, nil
}
