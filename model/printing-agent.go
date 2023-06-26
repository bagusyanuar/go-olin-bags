package model

import (
	"time"

	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	PrintingAgentTableName = "printing_agents"
)

type PrintingAgent struct {
	ID                uuid.UUID `json:"id"`
	ProductionHouseID uuid.UUID `json:"production_house_id"`
	Name              string    `json:"name"`
	Phone             string    `json:"phone"`
	Address           string    `json:"address"`
	common.WithTimestampsModel
	ProductionHouse *ProductionHouse `gorm:"foreignKey:ProductionHouseID" json:"production_house"`
}

func (printingAgent *PrintingAgent) BeforeCreate(tx *gorm.DB) (err error) {
	printingAgent.ID = uuid.New()
	printingAgent.CreatedAt = time.Now()
	printingAgent.UpdatedAt = time.Now()
	return
}

func (PrintingAgent) TableName() string {
	return PrintingAgentTableName
}
