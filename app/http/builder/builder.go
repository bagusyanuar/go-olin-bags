package builder

import (
	"github.com/bagusyanuar/go-olin-bags/app/config"
	"github.com/bagusyanuar/go-olin-bags/app/http/controller"
	adminCtrl "github.com/bagusyanuar/go-olin-bags/app/http/controller/admin"
	"github.com/bagusyanuar/go-olin-bags/app/repositories"
	adminRepo "github.com/bagusyanuar/go-olin-bags/app/repositories/admin"
	"github.com/bagusyanuar/go-olin-bags/app/service"
	adminSvc "github.com/bagusyanuar/go-olin-bags/app/service/admin"
	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/bagusyanuar/go-olin-bags/router"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Builder struct {
	DB                   *gorm.DB
	Config               *config.Config
	WelcomeController    *controller.WelcomeController
	AuthController       *controller.AuthController
	AgentAdminController *adminCtrl.AgentController
}

func NewBuilder(db *gorm.DB, cfg *config.Config) Builder {
	return Builder{
		DB:     db,
		Config: cfg,
	}
}

func (b *Builder) Build() {
	authRepository := repositories.NewAuthRepository(b.DB)
	agentRepository := adminRepo.NewAgentRepository(b.DB)

	authService := service.NewAuthService(b.Config.JWT, authRepository)
	agentService := adminSvc.NewAgentService(agentRepository)

	authController := controller.NewAuthController(authService)
	b.AuthController = &authController

	agentController := adminCtrl.NewAgentController(agentService)
	b.AgentAdminController = &agentController

	welcomeController := controller.NewWelcomeController(b.Config)
	b.WelcomeController = &welcomeController
}

func (b *Builder) CreateRoute(group *gin.RouterGroup) {
	routes := b.routers()
	for _, route := range routes {
		group.Handle(route.Method, route.Group+route.Path, route.Handler...)
	}
}

func (b *Builder) routers() []*common.Route {
	return router.Routers(
		b.WelcomeController,
		b.AuthController,
		b.AgentAdminController,
	)
}

// func (b *Builder) CreateRoute() []*common.Route {
// 	return router.Routers(
// 		b.WelcomeController,
// 	)
// }

// func CreateRoute(cfg *config.Config) []*common.Route {
// 	welcomeController := controller.NewWelcomeController(cfg)
// 	return router.Routers(
// 		&welcomeController,
// 	)
// }
