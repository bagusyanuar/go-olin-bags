package server

import (
	"fmt"
	"net/http"

	"github.com/bagusyanuar/go-olin-bags/app/http/builder"
	"github.com/bagusyanuar/go-olin-bags/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Serve(cfg *config.Config, db *gorm.DB) {
	//gin init
	server := gin.Default()
	server.Use(gin.Recovery())
	server.SetTrustedProxies([]string{"127.0.0.1", "localhost"})

	//set not found route response
	server.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "route not found",
			"data":    nil,
		})
	})

	//build up scheme & route
	b := builder.NewBuilder(db, cfg)
	b.Build(server)

	//listening server
	listen := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	server.Run(listen)
}
