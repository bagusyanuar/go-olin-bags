package controller

import (
	"github.com/bagusyanuar/go-olin-bags/config"
	"github.com/gin-gonic/gin"
)

type WelcomeController struct {
	Config *config.Config
}

func NewWelcomeController(cfg *config.Config) WelcomeController {
	return WelcomeController{
		Config: cfg,
	}
}

func (c *WelcomeController) Index(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"app_name":    c.Config.AppName,
		"app_version": c.Config.AppVersion,
	})
}
