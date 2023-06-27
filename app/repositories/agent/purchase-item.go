package agent

import (
	"github.com/bagusyanuar/go-olin-bags/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PurchaseItem interface {
	List(accountID string) ([]model.PurchaseItem, error)
	AddItem(entity model.PurchaseItem) (*model.PurchaseItem, error)
	DropItem(accountID string, id string) error
}

type PurchaseItemRepository struct {
	Database *gorm.DB
}

// AddItem implements PurchaseItem.
func (r *PurchaseItemRepository) AddItem(entity model.PurchaseItem) (*model.PurchaseItem, error) {
	if err := r.Database.
		Omit(clause.Associations).
		Create(&entity).
		Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// DropItem implements PurchaseItem.
func (r *PurchaseItemRepository) DropItem(accountID string, id string) error {
	panic("unimplemented")
}

// List implements PurchaseItem.
func (r *PurchaseItemRepository) List(accountID string) ([]model.PurchaseItem, error) {
	var data []model.PurchaseItem
	if err := r.Database.
		Where("purchasing_id IS NULL").
		Where("purchaser_id = ?", accountID).
		Preload("Purchasing").
		Preload("Item").
		Order("created_at ASC").
		Find(&data).
		Error; err != nil {
		return data, err
	}
	return data, nil
}

func NewPurchaseItemRepository(database *gorm.DB) PurchaseItem {
	return &PurchaseItemRepository{
		Database: database,
	}
}
