package model

import (
	"time"

	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type City struct {
	ID         uuid.UUID `json:"id"`
	ProvinceID uuid.UUID `json:"province_id"`
	Code       string    `json:"code"`
	Name       string    `json:"name"`
	Province   Province  `gorm:"foreignKey:ProvinceID" json:"province"`
	common.WithTimestampsModel
}

func (city *City) BeforeCreate(tx *gorm.DB) (err error) {
	city.ID = uuid.New()
	city.CreatedAt = time.Now()
	city.UpdatedAt = time.Now()
	return
}

func (City) TableName() string {
	return common.CityTableName
}
