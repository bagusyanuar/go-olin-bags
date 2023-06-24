package model

import (
	"time"

	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	AgentTableName = "agents"
)

type Agent struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	CityID    uuid.UUID `json:"city_id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Balance   float64   `json:"balance"`
	common.WithTimestampsModel
	User User `gorm:"foreignKey:UserID" json:"user"`
	City City `gorm:"foreignKey:CityID" json:"city"`
}

func (agent *Agent) BeforeCreate(tx *gorm.DB) (err error) {
	agent.ID = uuid.New()
	agent.CreatedAt = time.Now()
	agent.UpdatedAt = time.Now()
	return
}

func (Agent) TableName() string {
	return AgentTableName
}
