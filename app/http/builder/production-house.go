package builder

import (
	controller "github.com/bagusyanuar/go-olin-bags/app/http/controller/production-house"
	repository "github.com/bagusyanuar/go-olin-bags/app/repositories/production-house"
	service "github.com/bagusyanuar/go-olin-bags/app/service/production-house"
	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/bagusyanuar/go-olin-bags/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductionHouseBuilder struct {
	Database *gorm.DB
	Config   *config.Config
	APIGroup *gin.RouterGroup
}

func NewProductionHouseBuilder(
	database *gorm.DB,
	config *config.Config,
	apiGroup *gin.RouterGroup,
) AgentBuilder {
	return AgentBuilder{
		Database: database,
		Config:   config,
		APIGroup: apiGroup,
	}
}

func (b *ProductionHouseBuilder) BuildScheme() {
	authRepository := repository.NewAuthRepository(b.Database)

	authService := service.NewAuthService(authRepository, b.Config.JWT)

	authController := controller.NewAuthController(authService, b.APIGroup)

	controllers := []any{
		&authController,
	}
	common.RegisterRoutes(controllers...)
}
