package models

import (
	"errors"
	"time"

	"github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Stock struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" json:"stockID"`
	Type         string    `gorm:"size:10;index:idx_type_province,unique" json:"type"`
	Description  string    `gorm:"size:255;not null" json:"description"`
	Availability bool      `gorm:"type:bool;default:true" json:"availability"`
	Cost         float32   `gorm:"not null" json:"-"`
	Price        float32   `gorm:"not null" json:"price"`
	Province     string    `gorm:"size:15;index:idx_type_province,unique" json:"province"`
	Unit         int       `gorm:"not null" json:"-"`
	Slot         int       `gorm:"not null" json:"slot"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type CreateStockInput struct {
	Type        string  `json:"type" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Province    string  `json:"province" binding:"required"`
	Cost        float32 `json:"cost" binding:"required"`
	Price       float32 `json:"price"`
	Unit        int     `json:"unit" binding:"required"`
	Slot        int     `json:"slot"`
}

type UpdateStockInput struct {
	Type        string  `json:"type"`
	Description string  `json:"description"`
	Province    string  `json:"province"`
	Cost        float32 `json:"cost"`
	Price       float32 `json:"price"`
	Unit        int     `json:"unit"`
	Slot        int     `json:"slot"`
}

type UpdateSlotInput struct {
	StockID uuid.UUID `json:"slotID"`
	Slot    int       `json:"slot"`
}

func (stock *Stock) BeforeSave(scope *gorm.DB) error {
	stock.ID = uuid.NewV4()
	return nil
}

func (stock *Stock) Save() error {
	if err := DB.Create(&stock).Error; err != nil {
		return err
	}
	return nil
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

func (stock *Stock) UpdateSlot(updateStock UpdateSlotInput) (*Stock, error) {
	if err := DB.Where("id = ?", updateStock.StockID).First(&stock).Error; err != nil {
		return &Stock{}, err
	}
	stock.Slot = stock.Slot + updateStock.Slot
	if err := DB.Model(&stock).Update("slot", stock.Slot).Error; err != nil {
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
