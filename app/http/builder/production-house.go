package builder

import (
	controller "github.com/bagusyanuar/go-olin-bags/app/http/controller/production-house"
	"github.com/bagusyanuar/go-olin-bags/app/http/middleware"
	repository "github.com/bagusyanuar/go-olin-bags/app/repositories/production-house"
	service "github.com/bagusyanuar/go-olin-bags/app/service/production-house"
	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/bagusyanuar/go-olin-bags/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductionHouseBuilder struct {
	Database   *gorm.DB
	Config     *config.Config
	APIGroup   *gin.RouterGroup
	Middleware middleware.AuthMiddleware
}

func NewProductionHouseBuilder(
	database *gorm.DB,
	config *config.Config,
	apiGroup *gin.RouterGroup,
	middleware middleware.AuthMiddleware,
) ProductionHouseBuilder {
	return ProductionHouseBuilder{
		Database:   database,
		Config:     config,
		APIGroup:   apiGroup,
		Middleware: middleware,
	}
}

func (b *ProductionHouseBuilder) BuildScheme() {
	authRepository := repository.NewAuthRepository(b.Database)
	profileRepository := repository.NewProfileRepository(b.Database)
	sewingAgentRepository := repository.NewSewingAgentRepository(b.Database)
	printingAgentRepository := repository.NewPrintingAgentRepository(b.Database)

	authService := service.NewAuthService(authRepository, b.Config.JWT)
	sewingAgentService := service.NewSewingAgentService(sewingAgentRepository, profileRepository)
	printingAgentService := service.NewPrintingAgentService(printingAgentRepository, profileRepository)

	authController := controller.NewAuthController(authService, b.APIGroup)
	sewingAgentController := controller.NewSewingAgentController(sewingAgentService, b.APIGroup, b.Middleware)
	printingAgentController := controller.NewPrintingAgentController(printingAgentService, b.APIGroup, b.Middleware)

	controllers := []any{
		&authController,
		&sewingAgentController,
		&printingAgentController,
	}
	common.RegisterRoutes(controllers...)
}
