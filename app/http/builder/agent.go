package builder

import (
	controller "github.com/bagusyanuar/go-olin-bags/app/http/controller/agent"
	"github.com/bagusyanuar/go-olin-bags/app/http/middleware"
	repository "github.com/bagusyanuar/go-olin-bags/app/repositories/agent"
	service "github.com/bagusyanuar/go-olin-bags/app/service/agent"
	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/bagusyanuar/go-olin-bags/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AgentBuilder struct {
	Database   *gorm.DB
	Config     *config.Config
	APIGroup   *gin.RouterGroup
	Middleware middleware.AuthMiddleware
}

func NewAgentBuilder(
	database *gorm.DB,
	config *config.Config,
	apiGroup *gin.RouterGroup,
	middleware middleware.AuthMiddleware,
) AgentBuilder {
	return AgentBuilder{
		Database:   database,
		Config:     config,
		APIGroup:   apiGroup,
		Middleware: middleware,
	}
}

func (b *AgentBuilder) BuildScheme() {
	authRepository := repository.NewAuthRepository(b.Database)
	itemRepository := repository.NewItemRepository(b.Database)
	purchaseItemRepository := repository.NewPurchaseItemRepository(b.Database)

	authService := service.NewAuthService(authRepository, b.Config.JWT)
	purchaseItemService := service.NewPurchaseItemService(purchaseItemRepository, itemRepository)

	authController := controller.NewAuthController(authService, b.APIGroup)
	purchaseItemController := controller.NewPurchaseItemController(purchaseItemService, b.APIGroup, b.Middleware)

	controllers := []any{
		&authController,
		&purchaseItemController,
	}
	common.RegisterRoutes(controllers...)
}
