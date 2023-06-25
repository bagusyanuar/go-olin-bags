package productionhouse

import (
	"github.com/bagusyanuar/go-olin-bags/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SewingAgent interface {
	FindAll(authorizedID, q string, limit, offset int) ([]model.SewingAgent, error)
	FindByID(authorizedID, id string) (*model.SewingAgent, error)
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
func (r *SewingAgentRepository) FindAll(authorizedID string, q string, limit int, offset int) ([]model.SewingAgent, error) {
	var data []model.SewingAgent
	if err := r.Database.
		Where("production_house_id = ?", authorizedID).
		Where("name Like ?", "%"+q+"%").
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
func (r *SewingAgentRepository) FindByID(authorizedID string, id string) (*model.SewingAgent, error) {
	data := new(model.SewingAgent)
	if err := r.Database.
		Where("production_house_id = ?", authorizedID).
		Where("id = ?", id).
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
