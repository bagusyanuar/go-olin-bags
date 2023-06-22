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
	adminGroup := api.Group("/admin")

	publicBuilder := NewPublicBuilder()
	publicBuilder.BuildPublicSheme(b.DB, b.Config, api)

	adminBuilder := NewAdminBuilder()
	adminBuilder.BuildScheme(b.DB, b.Config, adminGroup)
}
