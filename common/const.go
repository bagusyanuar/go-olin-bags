package common

import (
	"time"

	"gorm.io/gorm"
)

const (
	UserTableName  = "users"
	AgentTableName = "agents"
)

type WithTimestampsModel struct {
	CreatedAt time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at"`
}
