package model

import (
	"time"

	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

const (
	PurchasingTableName = "purchasings"
)

type Purchasing struct {
	ID                  uuid.UUID      `json:"id"`
	PurchaserID         uuid.UUID      `json:"purchaser_id"`
	ProductionHouseID   uuid.UUID      `json:"production_house_id"`
	AccessorID          uuid.UUID      `json:"accessor_id"`
	PurchaseNumber      string         `json:"purchase_number"`
	Date                datatypes.Date `json:"date"`
	ShippingDestination string         `json:"shipping_destination"`
	SubTotal            uint64         `json:"sub_total"`
	ShippingCost        uint64         `json:"shipping_cost"`
	Discount            uint64         `json:"discount"`
	Total               uint64         `json:"total"`
	Status              uint8          `json:"status"`
	common.WithTimestampsModel
	Purchaser       *User            `gorm:"foreignKey:PurchaserID" json:"purchaser"`
	ProductionHouse *ProductionHouse `gorm:"foreignKey:ProductionHouseID" json:"production_house"`
	Accessor        *User            `gorm:"foreignKey:AccessorID" json:"accessor"`
	PurchaseItems   []PurchaseItem   `gorm:"foreignKey:PurchasingID" json:"purchase_item"`
}

func (purchasing *Purchasing) BeforeCreate(tx *gorm.DB) (err error) {
	purchasing.ID = uuid.New()
	purchasing.CreatedAt = time.Now()
	purchasing.UpdatedAt = time.Now()
	return
}

func (Purchasing) TableName() string {
	return PurchasingTableName
}
