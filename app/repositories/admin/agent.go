package admin

import (
	"github.com/bagusyanuar/go-olin-bags/model"
	"gorm.io/gorm"
)

type Agent interface {
	FindAll() ([]model.Agent, error)
}

type AgentRepository struct {
	Database *gorm.DB
}

// FindAll implements Agent.
func (r *AgentRepository) FindAll() ([]model.Agent, error) {
	var data []model.Agent
	if err := r.Database.Debug().Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func NewAgentRepository(db *gorm.DB) Agent {
	return &AgentRepository{
		Database: db,
	}
}
