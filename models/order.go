package models

import (
	"time"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"github.com/satori/go.uuid"
)

type Order struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key" json:"OrderID"`
	Quantity int `gorm:"not null" json:"Quantity"`
	Stocks []*Stock `gorm:"many2many:order_stocks;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"Stocks"`
	UserID uuid.UUID `gorm:"foreignKey:ID"`
	CreatedAt time.Time `json:"CreatedAt"`
  UpdatedAt time.Time `json:"UpdatedAt"`
}

type OrderMadeInput struct {
	Quantity int `json:"Quantity" binding:"required"`
	UserID uuid.UUID `json:"UserID" binding:"required"`
	Stocks []Stock
}

type OrderRequest struct {
	Quantity int
	Stock string
}

type OrderInputs []OrderRequest

type OrderInput struct {
	Stocks []OrderRequest `json:"Stocks" binding:"required"`
}

func (order *Order) BeforeSave(scope *gorm.DB) error {
	order.ID = uuid.NewV4()
	return nil
 }

 func (order *Order) Save() (*Order, error) {
	if err := DB.Omit("Stocks.*").Create(&order).Error; err != nil {
		return &Order{}, err
	}
	return order, nil
}

func (order *Order) Association() (*Order, error) {
	if er := DB.Association("Stocks").Append(&order.Stocks); er != nil {
		return &Order{}, er
	}
	return order, nil
}

func (order *Order) FindAll() []Order {
	var orders []Order
	DB.Preload(clause.Associations).Find(&orders)
	return orders
}

func (order *Order) FindOrderByID(ID string) (*Order, error) {
	if err := DB.Where("id = ?", ID).First(&order).Error; err != nil {
		return order, err
	}
	return order, nil
}
