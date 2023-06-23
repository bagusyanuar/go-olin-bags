package controller

import (
	"fmt"
	"net/http"

	"github.com/bagusyanuar/go-olin-bags/app/http/request"
	"github.com/bagusyanuar/go-olin-bags/app/service"
	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService service.Auth
	APIGroup    *gin.RouterGroup
}

func NewAuthController(authService service.Auth, apiGroup *gin.RouterGroup) AuthController {
	return AuthController{
		AuthService: authService,
		APIGroup:    apiGroup,
	}
}

func (c *AuthController) RegisterRoutes() {
	route := c.APIGroup.Group("/auth")
	{
		route.POST("/", c.SignIn)
	}
}

func (c *AuthController) SignIn(ctx *gin.Context) {
	var request request.CreateSignInRequest
	ctx.BindJSON(&request)

	token, err := c.AuthService.SignIn(request)
	if err != nil {
		ctx.JSON(500, common.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("internal server error (%s)", err.Error()),
			Data:    nil,
		})
		return
	}
	ctx.JSON(200, common.APIResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    token,
	})
}
