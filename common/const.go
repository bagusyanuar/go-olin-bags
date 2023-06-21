package common

import (
	"time"

	"gorm.io/gorm"
)

// table name constant
const (
	UserTableName            = "users"
	ProvinceTableName        = "provinces"
	CityTableName            = "cities"
	ProductionHouseTableName = "production_houses"
	AgentTableName           = "agents"
)

const (
	DefaultLimit int = 5
)

type WithTimestampsModel struct {
	CreatedAt time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at"`
}
