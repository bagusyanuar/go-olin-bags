package builder

import (
	controller "github.com/bagusyanuar/go-olin-bags/app/http/controller/admin"
	repository "github.com/bagusyanuar/go-olin-bags/app/repositories/admin"
	service "github.com/bagusyanuar/go-olin-bags/app/service/admin"
	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/bagusyanuar/go-olin-bags/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AdminBuilder struct {
	Database *gorm.DB
	Config   *config.Config
	APIGroup *gin.RouterGroup
}

func NewAdminBuilder(
	database *gorm.DB,
	config *config.Config,
	apiGroup *gin.RouterGroup,
) AdminBuilder {
	return AdminBuilder{
		Database: database,
		Config:   config,
		APIGroup: apiGroup,
	}
}

func (b *AdminBuilder) BuildScheme() {
	agentRepository := repository.NewAgentRepository(b.Database)
	provinceRepository := repository.NewProvinceReposiotry(b.Database)
	cityRepository := repository.NewCityRepository(b.Database)

	agentService := service.NewAgentService(agentRepository)
	provinceService := service.NewProvinceService(provinceRepository)
	cityService := service.NewCityService(cityRepository)

	agentController := controller.NewAgentController(agentService)
	provinceController := controller.NewProvinceController(provinceService, b.APIGroup)
	cityController := controller.NewCityController(cityService, b.APIGroup)

	controllers := []any{
		&agentController,
		&provinceController,
		&cityController,
	}
	common.RegisterRoutes(controllers...)
}
