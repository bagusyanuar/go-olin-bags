package model

import (
	"time"

	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	PurchaseItemTableName = "purchase_items"
)

type PurchaseItem struct {
	ID           uuid.UUID  `json:"id"`
	PurchasingID *uuid.UUID `json:"purchasing_id"`
	PurchaserID  uuid.UUID  `json:"purchaser_id"`
	ItemID       *uuid.UUID `json:"item_id"`
	Price        uint64     `json:"price"`
	Qty          uint32     `json:"qty"`
	Total        uint64     `json:"total"`
	common.WithTimestampsModel
	Purchasing *Purchasing `gorm:"foreignKey:PurchasingID" json:"Purchasing"`
	Purchaser  *User       `gorm:"foreignKey:PurchaserID" json:"purchaser"`
	Item       *Item       `gorm:"foreignKey:ItemID" json:"item"`
}

func (purchaseItem *PurchaseItem) BeforeCreate(tx *gorm.DB) (err error) {
	purchaseItem.ID = uuid.New()
	purchaseItem.CreatedAt = time.Now()
	purchaseItem.UpdatedAt = time.Now()
	return
}

func (PurchaseItem) TableName() string {
	return PurchaseItemTableName
}
