package agent

import (
	"github.com/bagusyanuar/go-olin-bags/model"
	"gorm.io/gorm"
)

type Item interface {
	GetItemByID(id string) (*model.Item, error)
}

type ItemRepository struct {
	Database *gorm.DB
}

// GetItemByID implements Item.
func (r *ItemRepository) GetItemByID(id string) (*model.Item, error) {
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
