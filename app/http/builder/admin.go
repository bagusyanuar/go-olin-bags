package builder

import (
	"github.com/bagusyanuar/go-olin-bags/app/config"
	controller "github.com/bagusyanuar/go-olin-bags/app/http/controller/admin"
	repository "github.com/bagusyanuar/go-olin-bags/app/repositories/admin"
	service "github.com/bagusyanuar/go-olin-bags/app/service/admin"
	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/bagusyanuar/go-olin-bags/router"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AdminBuilder struct {
	AgentController *controller.AgentController
}

func NewAdminBuilder() AdminBuilder {
	return AdminBuilder{}
}

func (b *AdminBuilder) BuildPublicSheme(db *gorm.DB, cfg *config.Config, group *gin.RouterGroup) {
	agentRepository := repository.NewAgentRepository(db)

	agentService := service.NewAgentService(agentRepository)

	agentController := controller.NewAgentController(agentService)
	b.AgentController = &agentController
	b.createRoutes(group)
}

func (b *AdminBuilder) routes() []*common.Route {
	return router.AdminRoutes(
		b.AgentController,
	)
}

func (b *AdminBuilder) createRoutes(group *gin.RouterGroup) {
	routes := b.routes()
	for _, route := range routes {
		group.Handle(route.Method, route.Group+route.Path, route.Handler)
	}
}
