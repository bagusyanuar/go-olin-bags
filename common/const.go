package common

import (
	"time"

	"gorm.io/gorm"
)

// table name constant
const (
	UserTableName = "users"

	ProductionHouseTableName = "production_houses"
)

const (
	DefaultLimit               int    = 5
	CodeDateFormat             string = "20060102150405"
	PurchasingOnWaiting        uint8  = 0
	PurchasingOnWaitingPayment uint8  = 1
	PurchasingOnPayment        uint8  = 2
	PurchasingOnProcess        uint8  = 3
)

type WithTimestampsModel struct {
	CreatedAt time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at"`
}
