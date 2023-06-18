package models

import (
	"time"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"github.com/satori/go.uuid"
)

type ShippingAddress struct {
	StreetNo int `json:"StreetNo" binding:"required,gte=0,lte=1000"`
	StreetName string  `json:"StreetName" binding:"required,max=50"`
	Province string  `json:"Province" binding:"required,max=50"`
	State string  `json:"State" binding:"required,max=50"`
}

type Order struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key" json:"OrderID"`
	TotalQuantity int `gorm:"not null" json:"TotalQuantity"`
	Products []Product `gorm:"many2many:order_products;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"Products"`
	UserID uuid.UUID `gorm:"foreignKey:ID"`
	User User `gorm:"belongs_to:user" json:"User"`
	ShippingAddress ShippingAddress `gorm:"embedded;embeddedPrefix:shippingaddress_"`
	CreatedAt time.Time `json:"CreatedAt"`
  UpdatedAt time.Time `json:"UpdatedAt"`
}

type OrderRequest struct {
	Quantity int `json:"Quantity" binding:"required,gte=1,lte=100"`
	StockID uuid.UUID `json:"StockID" binding:"required"`
}

type OrderInputs struct {
	Products []*OrderRequest `json:"Products" binding:"required"`
	ShippingAddress ShippingAddress `json:"ShippingAddress" binding:"required"`
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
	DB.Preload(clause.Associations).Preload("Products.Stock").Find(&orders)
	return orders
}

func (order *Order) FindOrderByID(ID string) (*Order, error) {
	if err := DB.Where("id = ?", ID).First(&order).Error; err != nil {
		return order, err
	}
	return order, nil
}
