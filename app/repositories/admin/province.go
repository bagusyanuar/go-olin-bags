package admin

import (
	"github.com/bagusyanuar/go-olin-bags/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// interface for province admin scheme
type Province interface {
	FindAll(q string, limit, offset int) ([]model.Province, error)
	FindByID(id string) (*model.Province, error)
	Create(entity model.Province) error
	Patch(id string, entity model.Province) error
	Delete(id string) error
}

// init province repository struct
type ProvinceRepository struct {
	Database *gorm.DB
}

// Delete implements Province.
func (r *ProvinceRepository) Delete(id string) error {
	if err := r.Database.Where("id = ?", id).Delete(&model.Province{}).Error; err != nil {
		return err
	}
	return nil
}

// Patch implements Province.
func (r *ProvinceRepository) Patch(id string, entity model.Province) error {
	if err := r.Database.
		Omit(clause.Associations).
		Where("id = ?", id).
		Updates(entity).
		Error; err != nil {
		return err
	}
	return nil
}

// Create implements Province.
func (r *ProvinceRepository) Create(entity model.Province) error {
	if err := r.Database.
		Omit(clause.Associations).
		Create(&entity).
		Error; err != nil {
		return err
	}
	return nil
}

// FindAll implements Province.
func (r *ProvinceRepository) FindAll(q string, limit, offset int) ([]model.Province, error) {
	var data []model.Province
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

// FindByID implements Province.
func (r *ProvinceRepository) FindByID(id string) (*model.Province, error) {
	data := new(model.Province)
	if err := r.Database.
		Where("id = ?", id).
		First(&data).
		Error; err != nil {
		return nil, err
	}
	return data, nil
}

func NewProvinceReposiotry(database *gorm.DB) Province {
	return &ProvinceRepository{
		Database: database,
	}
}
