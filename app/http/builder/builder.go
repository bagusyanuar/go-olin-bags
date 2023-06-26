package builder

import (
	. "github.com/bagusyanuar/go-olin-bags/app/http/middleware"
	"github.com/bagusyanuar/go-olin-bags/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Builder struct {
	DB     *gorm.DB
	Config *config.Config
}

func NewBuilder(db *gorm.DB, cfg *config.Config) Builder {
	return Builder{
		DB:     db,
		Config: cfg,
	}
}

func (b *Builder) Build(server *gin.Engine) {

	middleware := NewAuthMiddleware(&b.Config.JWT)

	api := server.Group("/api/v1")
	//endpoint group for public
	publicGroup := api.Group("/")

	//endpoint group for admin
	adminGroup := api.Group("/admin")
	// adminGroup.Use(middleware.IsAuth())

	//endpoint group for production house
	productionHouseGroup := api.Group("/production-house")
	// productionHouseGroup.Use(middleware.IsAuth())

	//endpoint group for agent with auth middleware
	agentGroup := api.Group("/agent")
	agentGroup.Use(middleware.IsAuth())

	publicBuilder := NewPublicBuilder(b.DB, b.Config, publicGroup)
	publicBuilder.BuildScheme()

	adminBuilder := NewAdminBuilder(b.DB, b.Config, adminGroup, middleware)
	adminBuilder.BuildScheme()

	productionHouseBuilder := NewProductionHouseBuilder(b.DB, b.Config, productionHouseGroup, middleware)
	productionHouseBuilder.BuildScheme()

	agentBuilder := NewAgentBuilder(b.DB, b.Config, agentGroup)
	agentBuilder.BuildScheme()
}
