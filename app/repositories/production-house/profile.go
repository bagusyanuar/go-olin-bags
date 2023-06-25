package productionhouse

import (
	"github.com/bagusyanuar/go-olin-bags/model"
	"gorm.io/gorm"
)

type Profile interface {
	GetProfile(authorizedID string) (*model.ProductionHouse, error)
}

type ProfileRepository struct {
	Database *gorm.DB
}

// GetProfile implements Profile.
func (r *ProfileRepository) GetProfile(authorizedID string) (*model.ProductionHouse, error) {
	data := new(model.ProductionHouse)
	if err := r.Database.
		Where("user_id = ?", authorizedID).
		Preload("User").
		Preload("City").
		First(&data).
		Error; err != nil {
		return nil, err
	}
	return data, nil
}

func NewProfileRepository(database *gorm.DB) Profile {
	return &ProfileRepository{
		Database: database,
	}
}
