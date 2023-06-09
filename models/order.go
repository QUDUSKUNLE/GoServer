package models

import (
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type ShippingAddress struct {
	StreetNo   int    `json:"streetNo" binding:"required,gte=0,lte=1000"`
	StreetName string `json:"streetName" binding:"required,max=50"`
	Province   string `json:"province" binding:"required,max=50"`
	State      string `json:"state" binding:"required,max=50"`
}

type Order struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key" json:"orderID"`
	TotalQuantity int       `gorm:"not null" json:"totalQuantity"`
	Products      []Product `gorm:"many2many:order_products;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"products"`
	ProfileID     uuid.UUID `gorm:"foreignKey:ID" json:"-"`
	Profile       Profile   `gorm:"belongs_to:user" json:"userProfile"`
	AddressID     uuid.UUID `gorm:"foreignKey:ID" json:"-"`
	Address       Address   `gorm:"belongs_to:address" json:"shippingAddress"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type OrderRequest struct {
	Quantity int       `json:"quantity" binding:"required,gte=1,lte=100"`
	StockID  uuid.UUID `json:"stockID" binding:"required"`
}

type OrderInputs struct {
	Products  []*OrderRequest `json:"products" binding:"required"`
	AddressID uuid.UUID       `json:"addressID" binding:"required"`
}

func (order *Order) BeforeSave(scope *gorm.DB) error {
	order.ID = uuid.NewV4()
	return nil
}

func (order *Order) Save() error {
	if err := DB.Create(&order).Error; err != nil {
		return err
	}
	return nil
}

func (order *Order) FindAll() []Order {
	orders := []Order{}
	DB.Preload(clause.Associations).Preload("Products.Stock").Find(&orders)
	return orders
}

func (order *Order) FindOrderByID(ID string) (*Order, error) {
	if err := DB.Where("id = ?", ID).Preload(clause.Associations).Preload("Products.Stock").First(&order).Error; err != nil {
		return order, err
	}
	return order, nil
}
