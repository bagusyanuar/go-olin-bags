package admin

import (
	"github.com/bagusyanuar/go-olin-bags/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type City interface {
	FindAll(q string, limit, offset int) ([]model.City, error)
	FindByID(id string) (*model.City, error)
	Create(entity model.City) error
}

type CityRepository struct {
	Database *gorm.DB
}

// Create implements City.
func (r *CityRepository) Create(entity model.City) error {
	if err := r.Database.
		Omit(clause.Associations).
		Create(&entity).
		Error; err != nil {
		return err
	}
	return nil
}

// FindAll implements City.
func (r *CityRepository) FindAll(q string, limit int, offset int) ([]model.City, error) {
	var data []model.City
	if err := r.Database.
		Where("name Like ?", "%"+q+"%").
		Preload("Province").
		Order("name ASC").
		Limit(limit).
		Offset(offset).
		Find(&data).
		Error; err != nil {
		return data, err
	}
	return data, nil
}

// FindByID implements City.
func (r *CityRepository) FindByID(id string) (*model.City, error) {
	data := new(model.City)
	if err := r.Database.
		Where("id = ?", id).
		First(&data).
		Error; err != nil {
		return nil, err
	}
	return data, nil
}

func NewCityRepository(database *gorm.DB) City {
	return &CityRepository{
		Database: database,
	}
}
