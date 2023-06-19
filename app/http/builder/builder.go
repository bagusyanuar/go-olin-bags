package builder

import (
	"github.com/bagusyanuar/go-olin-bags/app/config"
	"github.com/bagusyanuar/go-olin-bags/app/http/controller"
	adminController "github.com/bagusyanuar/go-olin-bags/app/http/controller/admin"
	adminRepo "github.com/bagusyanuar/go-olin-bags/app/repositories/admin"
	adminSvc "github.com/bagusyanuar/go-olin-bags/app/service/admin"
	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/bagusyanuar/go-olin-bags/router"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Builder struct {
	DB            *gorm.DB
	Config        *config.Config
	PublicBuilder publicBuilder
	AdminBuilder  adminBuilder
	// WelcomeController    *controller.WelcomeController
	// AuthController       *controller.AuthController
	// AgentAdminController *adminController.AgentController
}

type publicBuilder struct {
	WelcomeController *controller.WelcomeController
	AuthController    *controller.AuthController
}

type adminBuilder struct {
	AgentController *adminController.AgentController
}

func NewBuilder(db *gorm.DB, cfg *config.Config) Builder {
	return Builder{
		DB:     db,
		Config: cfg,
	}
}

func (b *Builder) Build(server *gin.Engine) {
	api := server.Group("/api/v1")

	publicBuilder := NewPublicBuilder()
	publicBuilder.BuildPublicSheme(b.DB, b.Config, api)
	// b.buildPublicScheme()
	// b.buildAdminScheme()
	// b.createRoute(api)
}

func (b *Builder) buildAdminScheme() {
	agentRepository := adminRepo.NewAgentRepository(b.DB)

	agentService := adminSvc.NewAgentService(agentRepository)

	agentController := adminController.NewAgentController(agentService)
	b.AdminBuilder.AgentController = &agentController
}

func (b *Builder) createRoute(group *gin.RouterGroup) {
	//creating public routes
	publicRoutes := b.createPublicRoutes()
	for _, publicRoute := range publicRoutes {
		group.Handle(publicRoute.Method, publicRoute.Group+publicRoute.Path, publicRoute.Handler)
	}

	//creating admin routes
	adminRoutes := b.createAdminRoutes()
	apiAdmin := group.Group("/admin")
	for _, adminRoute := range adminRoutes {
		apiAdmin.Handle(adminRoute.Method, adminRoute.Group+adminRoute.Path, adminRoute.Handler)
	}
}

func (b *Builder) createPublicRoutes() []*common.Route {
	return router.PublicRoutes(
		b.PublicBuilder.WelcomeController,
		b.PublicBuilder.AuthController,
	)
}

func (b *Builder) createAdminRoutes() []*common.Route {
	return router.AdminRoutes(
		b.AdminBuilder.AgentController,
	)
}
