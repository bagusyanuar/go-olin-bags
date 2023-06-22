package builder

import (
	controller "github.com/bagusyanuar/go-olin-bags/app/http/controller/admin"
	repository "github.com/bagusyanuar/go-olin-bags/app/repositories/admin"
	service "github.com/bagusyanuar/go-olin-bags/app/service/admin"
	"github.com/bagusyanuar/go-olin-bags/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AdminBuilder struct {
	AgentController    *controller.AgentController
	ProvinceController *controller.ProvinceController
}

func NewAdminBuilder() AdminBuilder {
	return AdminBuilder{}
}

func (b *AdminBuilder) BuildScheme(db *gorm.DB, cfg *config.Config, endpointGroup *gin.RouterGroup) {
	agentRepository := repository.NewAgentRepository(db)
	provinceRepository := repository.NewProvinceReposiotry(db)

	agentService := service.NewAgentService(agentRepository)
	provinceService := service.NewProvinceService(provinceRepository)

	agentController := controller.NewAgentController(agentService)
	b.AgentController = &agentController
	provinceController := controller.NewProvinceController(provinceService)
	provinceController.RegisterRoutes(endpointGroup)

	// cityController := controller.NewCityController()
	// b.ProvinceController = &provinceController
	// b.createRoutes(group)
}

// func (b *AdminBuilder) routes() []*common.Route {
// 	return router.AdminRoutes(
// 		b.AgentController,
// 		b.ProvinceController,
// 	)
// }

// func (b *AdminBuilder) createRoutes(group *gin.RouterGroup) {
// 	routes := b.routes()
// 	for _, route := range routes {
// 		group.Handle(route.Method, route.Group+route.Path, route.Handler)
// 	}
// }
