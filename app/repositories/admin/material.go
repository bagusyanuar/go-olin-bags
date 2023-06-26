package admin

import (
	"github.com/bagusyanuar/go-olin-bags/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Material interface {
	FindAll(q string, limit, offset int) ([]model.Material, error)
	FindByID(id string) (*model.Material, error)
	Create(entity model.Material) (*model.Material, error)
}

type MaterialRepository struct {
	Database *gorm.DB
}

// Create implements Material.
func (r *MaterialRepository) Create(entity model.Material) (*model.Material, error) {
	if err := r.Database.
		Omit(clause.Associations).
		Create(&entity).
		Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// FindAll implements Material.
func (r *MaterialRepository) FindAll(q string, limit int, offset int) ([]model.Material, error) {
	var data []model.Material
	if err := r.Database.
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

// FindByID implements Material.
func (r *MaterialRepository) FindByID(id string) (*model.Material, error) {
	data := new(model.Material)
	if err := r.Database.
		Where("id = ?", id).
		First(&data).
		Error; err != nil {
		return nil, err
	}
	return data, nil
}

func NewMaterialRepository(database *gorm.DB) Material {
	return &MaterialRepository{
		Database: database,
	}
}
