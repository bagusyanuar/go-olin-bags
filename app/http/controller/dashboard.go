package controller

import "github.com/gin-gonic/gin"

type DashboardController struct{}

func NewDashboardController() DashboardController {
	return DashboardController{}
}

func (c *DashboardController) Index(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "dashbooard",
	})
}
