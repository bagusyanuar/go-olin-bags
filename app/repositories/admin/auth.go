package admin

import (
	"github.com/bagusyanuar/go-olin-bags/model"
	"gorm.io/gorm"
)

type Auth interface {
	SignIn(entity model.User) (*model.User, error)
}

type AuthRepository struct {
	Database *gorm.DB
}

// SignIn implements Auth.
func (r *AuthRepository) SignIn(entity model.User) (*model.User, error) {
	user := new(model.User)
	if err := r.Database.
		Where("JSON_SEARCH(roles, 'all', 'administrator') IS NOT NULL").
		Where("username = ?", entity.Username).
		First(&user).
		Error; err != nil {
		return nil, err
	}
	return user, nil
}

func NewAuthRepository(database *gorm.DB) Auth {
	return &AuthRepository{
		Database: database,
	}
}
