package builder

import (
	"github.com/bagusyanuar/go-olin-bags/app/config"
	"github.com/bagusyanuar/go-olin-bags/app/http/controller"
	"github.com/bagusyanuar/go-olin-bags/app/http/service"
	"github.com/bagusyanuar/go-olin-bags/app/repositories"
	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/bagusyanuar/go-olin-bags/router"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Builder struct {
	DB                *gorm.DB
	Config            *config.Config
	WelcomeController *controller.WelcomeController
	AuthController    *controller.AuthController
}

func NewBuilder(db *gorm.DB, cfg *config.Config) Builder {
	return Builder{
		DB:     db,
		Config: cfg,
	}
}

func (b *Builder) Build() {
	authRepository := repositories.NewAuthRepository(b.DB)

	authService := service.NewAuthService(authRepository)
	authController := controller.NewAuthController(authService)
	b.AuthController = &authController

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
