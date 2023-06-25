package builder

import (
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
	api := server.Group("/api/v1")
	publicGroup := api.Group("/")
	adminGroup := api.Group("/admin")
	agentGroup := api.Group("/agent")

	publicBuilder := NewPublicBuilder(b.DB, b.Config, publicGroup)
	publicBuilder.BuildScheme()

	adminBuilder := NewAdminBuilder(b.DB, b.Config, adminGroup)
	adminBuilder.BuildScheme()

	agentBuilder := NewAgentBuilder(b.DB, b.Config, agentGroup)
	agentBuilder.BuildScheme()
}
