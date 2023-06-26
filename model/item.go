package model

import (
	"time"

	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	ItemTableName = "items"
)

type Item struct {
	ID          uuid.UUID `json:"id"`
	MaterialID  uuid.UUID `json:"material_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int64     `json:"price"`
	common.WithTimestampsModel
	Material *Material `gorm:"foreignKey:MaterialID" json:"material"`
}

func (item *Item) BeforeCreate(tx *gorm.DB) (err error) {
	item.ID = uuid.New()
	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()
	return
}

func (Item) TableName() string {
	return ItemTableName
}
