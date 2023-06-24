package admin

import (
	"github.com/bagusyanuar/go-olin-bags/model"
	"gorm.io/gorm"
)

type ProductionHouse interface {
	FindAll(q string, limit, offset int) ([]model.ProductionHouse, error)
	FindByID(id string) (*model.ProductionHouse, error)
	Create(entity model.ProductionHouse) (*model.ProductionHouse, error)
}

type ProductionHouseRepository struct {
	Database *gorm.DB
}

// Create implements ProductionHouse.
func (r *ProductionHouseRepository) Create(entity model.ProductionHouse) (*model.ProductionHouse, error) {
	tx := r.Database.Begin()
	defer func() {
		if rcr := recover(); rcr != nil {
			tx.Rollback()
			return
		}
	}()

	// if err := tx.Omit(clause.Associations).
	// 	Create(&entityUser).
	// 	Error; err != nil {
	// 	return err
	// }

	// user := entityUser
	// entityProductionHouse.UserID = user.ID
	if err := tx.Omit("City").
		Create(&entity).
		Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &entity, nil
}

// FindAll implements ProductionHouse.
func (r *ProductionHouseRepository) FindAll(q string, limit int, offset int) ([]model.ProductionHouse, error) {
	var data []model.ProductionHouse
	if err := r.Database.
		Where("name Like ?", "%"+q+"%").
		Preload("City").
		Order("name ASC").
		Limit(limit).
		Offset(offset).
		Find(&data).
		Error; err != nil {
		return data, err
	}
	return data, nil
}

// FindByID implements ProductionHouse.
func (r *ProductionHouseRepository) FindByID(id string) (*model.ProductionHouse, error) {
	data := new(model.ProductionHouse)
	if err := r.Database.
		Where("id = ?", id).
		First(&data).
		Error; err != nil {
		return nil, err
	}
	return data, nil
}

func NewProductionHouseRepository(database *gorm.DB) ProductionHouse {
	return &ProductionHouseRepository{
		Database: database,
	}
}
