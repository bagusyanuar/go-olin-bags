package admin

import (
	"github.com/bagusyanuar/go-olin-bags/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Item interface {
	FindAll(q string, limit, offset int) ([]model.Item, error)
	FindByID(id string) (*model.Item, error)
	Create(entity model.Item) (*model.Item, error)
}

type ItemRepository struct {
	Database *gorm.DB
}

// Create implements Item.
func (r *ItemRepository) Create(entity model.Item) (*model.Item, error) {
	if err := r.Database.
		Omit(clause.Associations).
		Create(&entity).
		Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// FindAll implements Item.
func (r *ItemRepository) FindAll(q string, limit int, offset int) ([]model.Item, error) {
	var data []model.Item
	if err := r.Database.
		Where("name Like ?", "%"+q+"%").
		Preload("Material").
		Order("name ASC").
		Limit(limit).
		Offset(offset).
		Find(&data).
		Error; err != nil {
		return data, err
	}
	return data, nil
}

// FindByID implements Item.
func (r *ItemRepository) FindByID(id string) (*model.Item, error) {
	data := new(model.Item)
	if err := r.Database.
		Where("id = ?", id).
		Preload("Material").
		First(&data).
		Error; err != nil {
		return nil, err
	}
	return data, nil
}

func NewItemRepository(database *gorm.DB) Item {
	return &ItemRepository{
		Database: database,
	}
}
