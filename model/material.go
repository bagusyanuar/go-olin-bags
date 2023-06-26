package model

import (
	"time"

	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	MaterialTableName = "materials"
)

type Material struct {
	ID   uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	Name string    `gorm:"type:varchar(255);not null;" json:"name"`
	common.WithTimestampsModel
}

func (material *Material) BeforeCreate(tx *gorm.DB) (err error) {
	material.ID = uuid.New()
	material.CreatedAt = time.Now()
	material.UpdatedAt = time.Now()
	return
}

func (Material) TableName() string {
	return MaterialTableName
}
