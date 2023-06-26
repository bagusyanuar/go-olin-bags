package model

import (
	"time"

	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductionHouse struct {
	ID        uuid.UUID `json:"id"`
	CityID    uuid.UUID `json:"city_id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	common.WithTimestampsModel
	City *City `gorm:"foreignKey:CityID" json:"city"`
}

func (productionHouse *ProductionHouse) BeforeCreate(tx *gorm.DB) (err error) {
	productionHouse.ID = uuid.New()
	productionHouse.CreatedAt = time.Now()
	productionHouse.UpdatedAt = time.Now()
	return
}

func (ProductionHouse) TableName() string {
	return common.ProductionHouseTableName
}
