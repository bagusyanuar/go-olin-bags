package builder

import (
	controller "github.com/bagusyanuar/go-olin-bags/app/http/controller/admin"
	"github.com/bagusyanuar/go-olin-bags/app/http/middleware"
	repository "github.com/bagusyanuar/go-olin-bags/app/repositories/admin"
	service "github.com/bagusyanuar/go-olin-bags/app/service/admin"
	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/bagusyanuar/go-olin-bags/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AdminBuilder struct {
	Database   *gorm.DB
	Config     *config.Config
	APIGroup   *gin.RouterGroup
	Middleware middleware.AuthMiddleware
}

func NewAdminBuilder(
	database *gorm.DB,
	config *config.Config,
	apiGroup *gin.RouterGroup,
	middleware middleware.AuthMiddleware,
) AdminBuilder {
	return AdminBuilder{
		Database:   database,
		Config:     config,
		APIGroup:   apiGroup,
		Middleware: middleware,
	}
}

func (b *AdminBuilder) BuildScheme() {
	authRepository := repository.NewAuthRepository(b.Database)
	agentRepository := repository.NewAgentRepository(b.Database)
	provinceRepository := repository.NewProvinceReposiotry(b.Database)
	cityRepository := repository.NewCityRepository(b.Database)
	productionHouseRepository := repository.NewProductionHouseRepository(b.Database)
	sewingAgentRepository := repository.NewSewingAgentRepository(b.Database)
	printingAgentRepository := repository.NewPrintingAgentRepository(b.Database)
	materialRepository := repository.NewMaterialRepository(b.Database)
	itemRepository := repository.NewItemRepository(b.Database)

	authService := service.NewAuthService(authRepository, b.Config.JWT)
	agentService := service.NewAgentService(agentRepository)
	provinceService := service.NewProvinceService(provinceRepository)
	cityService := service.NewCityService(cityRepository)
	productionHouseService := service.NewProductionHouseService(productionHouseRepository)
	sewingAgentService := service.NewSewingAgentService(sewingAgentRepository)
	printingAgentService := service.NewPrintingAgentService(printingAgentRepository)
	materialService := service.NewMaterialService(materialRepository)
	itemService := service.NewItemService(itemRepository)

	authController := controller.NewAuthController(authService, b.APIGroup)
	agentController := controller.NewAgentController(agentService, b.APIGroup)
	provinceController := controller.NewProvinceController(provinceService, b.APIGroup)
	cityController := controller.NewCityController(cityService, b.APIGroup)
	productionHouseController := controller.NewProductionHouseController(productionHouseService, b.APIGroup)
	sewingAgentController := controller.NewSewingAgentController(sewingAgentService, b.APIGroup, b.Middleware)
	printingAgentController := controller.NewPrintingAgentController(printingAgentService, b.APIGroup, b.Middleware)
	materialController := controller.NewMaterialController(materialService, b.APIGroup, b.Middleware)
	itemController := controller.NewItemController(itemService, b.APIGroup, b.Middleware)

	controllers := []any{
		&authController,
		&agentController,
		&provinceController,
		&cityController,
		&productionHouseController,
		&sewingAgentController,
		&printingAgentController,
		&materialController,
		&itemController,
	}
	common.RegisterRoutes(controllers...)
}
