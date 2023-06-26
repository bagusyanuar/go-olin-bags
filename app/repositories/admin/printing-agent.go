package admin

import (
	"github.com/bagusyanuar/go-olin-bags/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PrintingAgent interface {
	FindAll(q string, limit, offset int) ([]model.PrintingAgent, error)
	FindByID(id string) (*model.PrintingAgent, error)
	Create(entity model.PrintingAgent) (*model.PrintingAgent, error)
}

type PrintingAgentRepository struct {
	Database *gorm.DB
}

// Create implements PrintingAgent.
func (r *PrintingAgentRepository) Create(entity model.PrintingAgent) (*model.PrintingAgent, error) {
	if err := r.Database.
		Omit(clause.Associations).
		Create(&entity).
		Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// FindAll implements PrintingAgent.
func (r *PrintingAgentRepository) FindAll(q string, limit int, offset int) ([]model.PrintingAgent, error) {
	var data []model.PrintingAgent
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

// FindByID implements PrintingAgent.
func (r *PrintingAgentRepository) FindByID(id string) (*model.PrintingAgent, error) {
	data := new(model.PrintingAgent)
	if err := r.Database.
		Where("id = ?", id).
		Preload("ProductionHouse").
		First(&data).
		Error; err != nil {
		return nil, err
	}
	return data, nil
}

func NewPrintingAgentRepository(database *gorm.DB) PrintingAgent {
	return &PrintingAgentRepository{
		Database: database,
	}
}
