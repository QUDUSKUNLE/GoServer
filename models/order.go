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
	Products []Product `gorm:"many2many:order_products;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"Products"`
	UserID uuid.UUID `gorm:"foreignKey:ID"`
	CreatedAt time.Time `json:"CreatedAt"`
  UpdatedAt time.Time `json:"UpdatedAt"`
}

type OrderRequest struct {
	Quantity int
	StockID uuid.UUID
}

type OrderInputs struct {
	Products []OrderRequest `json:"Stocks" binding:"required"`
}

func (order *Order) BeforeSave(scope *gorm.DB) error {
	order.ID = uuid.NewV4()
	return nil
 }

 func (order *Order) Save() (*Order, error) {
	if err := DB.Create(&order).Error; err != nil {
		return &Order{}, err
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
