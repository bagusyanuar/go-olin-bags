package admin

import (
	"github.com/bagusyanuar/go-olin-bags/model"
	"gorm.io/gorm"
)

type City interface {
	FindAll(param string, limit, offset int) ([]model.City, error)
	FindByID(id string) (*model.City, error)
	Create(entity model.City) error
}

type CityRepository struct {
	Database *gorm.DB
}

// Create implements City.
func (r *CityRepository) Create(entity model.City) error {
	panic("unimplemented")
}

// FindAll implements City.
func (r *CityRepository) FindAll(param string, limit int, offset int) ([]model.City, error) {
	panic("unimplemented")
}

// FindByID implements City.
func (r *CityRepository) FindByID(id string) (*model.City, error) {
	panic("unimplemented")
}

func NewCityRepository(database *gorm.DB) City {
	return &CityRepository{
		Database: database,
	}
}
