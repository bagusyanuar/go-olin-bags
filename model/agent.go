package model

import (
	"time"

	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Agent struct {
	ID      uuid.UUID `json:"id"`
	UserID  uuid.UUID `json:"user_id"`
	Name    string    `json:"name"`
	Phone   string    `json:"phone"`
	Address string    `json:"address"`
	common.WithTimestampsModel
	User User `gorm:"foreignKey:UserID" json:"user"`
}

func (agent *Agent) BeforeCreate(tx *gorm.DB) (err error) {
	agent.ID = uuid.New()
	agent.CreatedAt = time.Now()
	agent.UpdatedAt = time.Now()
	return
}

func (Agent) TableName() string {
	return common.AgentTableName
}
