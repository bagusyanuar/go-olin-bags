package model

import (
	"time"

	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Province struct {
	ID   uuid.UUID `json:"id"`
	Code string    `json:"code"`
	Name string    `json:"name"`
	common.WithTimestampsModel
}

func (province *Province) BeforeCreate(tx *gorm.DB) (err error) {
	province.ID = uuid.New()
	province.CreatedAt = time.Now()
	province.UpdatedAt = time.Now()
	return
}

func (Province) TableName() string {
	return common.ProvinceTableName
}
