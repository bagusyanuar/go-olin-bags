package admin

import (
	"github.com/bagusyanuar/go-olin-bags/model"
	"gorm.io/gorm"
)

// interface for province admin scheme
type Province interface {
	FindAll() ([]model.Province, error)
	FindByID() (*model.Province, error)
	Create(entity model.Province) error
}

// init province repository struct
type ProvinceRepository struct {
	Database *gorm.DB
}

// Create implements Province.
func (r *ProvinceRepository) Create(entity model.Province) error {
	panic("unimplemented")
}

// FindAll implements Province.
func (r *ProvinceRepository) FindAll() ([]model.Province, error) {
	panic("unimplemented")
}

// FindByID implements Province.
func (r *ProvinceRepository) FindByID() (*model.Province, error) {
	panic("unimplemented")
}

func NewProvinceReposiotry(database *gorm.DB) Province {
	return &ProvinceRepository{
		Database: database,
	}
}
