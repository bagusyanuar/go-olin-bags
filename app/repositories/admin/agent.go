package admin

import (
	"github.com/bagusyanuar/go-olin-bags/model"
	"gorm.io/gorm"
)

type Agent interface {
	FindAll() ([]model.Agent, error)
	FindByID(id string) (*model.Agent, error)
	Create(entity model.Agent) (*model.Agent, error)
}

type AgentRepository struct {
	Database *gorm.DB
}

// Create implements Agent.
func (r *AgentRepository) Create(entity model.Agent) (*model.Agent, error) {
	tx := r.Database.Begin()
	defer func() {
		if recover := recover(); recover != nil {
			tx.Rollback()
			return
		}
	}()
	if err := tx.Omit("City").
		Create(&entity).
		Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &entity, nil
}

// FindByID implements Agent.
func (r *AgentRepository) FindByID(id string) (*model.Agent, error) {
	var data *model.Agent
	if err := r.Database.Debug().First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
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
