package admin

import (
	"github.com/bagusyanuar/go-olin-bags/model"
	"gorm.io/gorm"
)

type ProductionHouse interface {
	FindAll(param string, limit, offset int) ([]model.ProductionHouse, error)
	FindByID(id string) (*model.ProductionHouse, error)
	Create(entity model.ProductionHouse) error
}

type ProductionHouseRepository struct {
	Database *gorm.DB
}

// Create implements ProductionHouse.
func (r *ProductionHouseRepository) Create(entity model.ProductionHouse) error {
	panic("unimplemented")
}

// FindAll implements ProductionHouse.
func (r *ProductionHouseRepository) FindAll(param string, limit int, offset int) ([]model.ProductionHouse, error) {
	panic("unimplemented")
}

// FindByID implements ProductionHouse.
func (r *ProductionHouseRepository) FindByID(id string) (*model.ProductionHouse, error) {
	panic("unimplemented")
}

func NewProductionHouseRepository(database *gorm.DB) ProductionHouse {
	return &ProductionHouseRepository{
		Database: database,
	}
}
