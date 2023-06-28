package agent

import (
	"github.com/bagusyanuar/go-olin-bags/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Purchasing interface {
	Checkout(accountID string, entity model.Purchasing, carts []model.PurchaseItem) (*model.Purchasing, error)
	List(accountID string, limit, offset int) ([]model.Purchasing, error)
}

type PurchasingRepository struct {
	Database *gorm.DB
}

// Checkout implements Purchasing.
func (r *PurchasingRepository) Checkout(accountID string, entity model.Purchasing, carts []model.PurchaseItem) (*model.Purchasing, error) {
	tx := r.Database.Begin()
	defer func() {
		if recover := recover(); recover != nil {
			tx.Rollback()
			return
		}
	}()
	if err := tx.
		Omit(clause.Associations).
		Create(&entity).
		Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	purchasingID := entity.ID

	if err := tx.Model(&carts).Update("purchasing_id", purchasingID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &entity, nil
}

// List implements Purchasing.
func (r *PurchasingRepository) List(accountID string, limit int, offset int) ([]model.Purchasing, error) {
	var data []model.Purchasing
	if err := r.Database.
		Where("purchaser_id = ?", accountID).
		Preload("ProductionHouse").
		Preload("PurchaseItems").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&data).
		Error; err != nil {
		return data, err
	}
	return data, nil
}

func NewPurchasingRepository(database *gorm.DB) Purchasing {
	return &PurchasingRepository{
		Database: database,
	}
}
