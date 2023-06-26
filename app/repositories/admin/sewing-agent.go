package admin

import (
	"github.com/bagusyanuar/go-olin-bags/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SewingAgent interface {
	FindAll(q string, limit, offset int) ([]model.SewingAgent, error)
	FindByID(id string) (*model.SewingAgent, error)
	Create(entity model.SewingAgent) (*model.SewingAgent, error)
}

type SewingAgentRepository struct {
	Database *gorm.DB
}

// Create implements SewingAgent.
func (r *SewingAgentRepository) Create(entity model.SewingAgent) (*model.SewingAgent, error) {
	if err := r.Database.
		Omit(clause.Associations).
		Create(&entity).
		Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// FindAll implements SewingAgent.
func (r *SewingAgentRepository) FindAll(q string, limit int, offset int) ([]model.SewingAgent, error) {
	var data []model.SewingAgent
	if err := r.Database.
		Where("name Like ?", "%"+q+"%").
		Preload("ProductionHouse").
		Order("name ASC").
		Limit(limit).
		Offset(offset).
		Find(&data).
		Error; err != nil {
		return data, err
	}
	return data, nil
}

// FindByID implements SewingAgent.
func (r *SewingAgentRepository) FindByID(id string) (*model.SewingAgent, error) {
	data := new(model.SewingAgent)
	if err := r.Database.
		Where("id = ?", id).
		Preload("ProductionHouse").
		First(&data).
		Error; err != nil {
		return nil, err
	}
	return data, nil
}

func NewSewingAgentRepository(database *gorm.DB) SewingAgent {
	return &SewingAgentRepository{
		Database: database,
	}
}
