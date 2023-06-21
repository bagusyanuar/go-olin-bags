package builder

import (
	"github.com/bagusyanuar/go-olin-bags/app/http/controller"
	"github.com/bagusyanuar/go-olin-bags/app/repositories"
	"github.com/bagusyanuar/go-olin-bags/app/service"
	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/bagusyanuar/go-olin-bags/config"
	"github.com/bagusyanuar/go-olin-bags/router"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PublicBuilder struct {
	WelcomeController *controller.WelcomeController
	AuthController    *controller.AuthController
}

func NewPublicBuilder() PublicBuilder {
	return PublicBuilder{}
}
func (b *PublicBuilder) BuildPublicSheme(db *gorm.DB, cfg *config.Config, group *gin.RouterGroup) {
	authRepository := repositories.NewAuthRepository(db)
	authservice := service.NewAuthService(authRepository, cfg.JWT)
	authController := controller.NewAuthController(authservice)
	b.AuthController = &authController
	b.createRoutes(group)
}

func (b *PublicBuilder) createRoutes(group *gin.RouterGroup) {
	routes := b.routes()
	for _, route := range routes {
		group.Handle(route.Method, route.Group+route.Path, route.Handler)
	}
}

func (b *PublicBuilder) routes() []*common.Route {
	return router.PublicRoutes(
		b.WelcomeController,
		b.AuthController,
	)
}
