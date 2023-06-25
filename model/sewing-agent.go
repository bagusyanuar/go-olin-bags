package model

import (
	"time"

	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	SewingAgentTableName = "sewing_agents"
)

type SewingAgent struct {
	ID                uuid.UUID `json:"id"`
	ProductionHouseID uuid.UUID `json:"production_house_id"`
	Name              string    `json:"name"`
	Phone             string    `json:"phone"`
	Address           string    `json:"address"`
	common.WithTimestampsModel
	ProductionHouse *ProductionHouse `gorm:"foreignKey:ProductionHouseID" json:"production_house"`
}

func (sewingAgent *SewingAgent) BeforeCreate(tx *gorm.DB) (err error) {
	sewingAgent.ID = uuid.New()
	sewingAgent.CreatedAt = time.Now()
	sewingAgent.UpdatedAt = time.Now()
	return
}

func (SewingAgent) TableName() string {
	return SewingAgentTableName
}
