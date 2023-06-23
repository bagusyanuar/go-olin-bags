package builder

import (
	"github.com/bagusyanuar/go-olin-bags/app/http/controller"
	"github.com/bagusyanuar/go-olin-bags/app/repositories"
	"github.com/bagusyanuar/go-olin-bags/app/service"
	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/bagusyanuar/go-olin-bags/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PublicBuilder struct {
	Database *gorm.DB
	Config   *config.Config
	APIGroup *gin.RouterGroup
}

func NewPublicBuilder(
	database *gorm.DB,
	config *config.Config,
	apiGroup *gin.RouterGroup,
) PublicBuilder {
	return PublicBuilder{
		Database: database,
		Config:   config,
		APIGroup: apiGroup,
	}
}
func (b *PublicBuilder) BuildScheme() {
	authRepository := repositories.NewAuthRepository(b.Database)
	authservice := service.NewAuthService(authRepository, b.Config.JWT)
	authController := controller.NewAuthController(authservice, b.APIGroup)

	controllers := []any{
		&authController,
	}
	common.RegisterRoutes(controllers...)
}
