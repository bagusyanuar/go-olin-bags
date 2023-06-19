package repositories

import (
	"github.com/bagusyanuar/go-olin-bags/model"
	"gorm.io/gorm"
)

type Auth interface {
	SignIn(u model.User) (*model.User, error)
}

type AuthRepository struct {
	Database *gorm.DB
}

// SignIn implements Auth.
func (r *AuthRepository) SignIn(u model.User) (*model.User, error) {
	user := new(model.User)
	if err := r.Database.Debug().Where("username = ?", u.Username).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func NewAuthRepository(db *gorm.DB) Auth {
	return &AuthRepository{
		Database: db,
	}
}
