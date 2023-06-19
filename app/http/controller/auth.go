package controller

import (
	"github.com/bagusyanuar/go-olin-bags/app/http/request"
	"github.com/bagusyanuar/go-olin-bags/app/http/service"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService service.Auth
}

func NewAuthController(authService service.Auth) AuthController {
	return AuthController{AuthService: authService}
}

func (c *AuthController) SignIn(ctx *gin.Context) {
	var request request.CreateSignInRequest
	ctx.BindJSON(&request)

	token, err := c.AuthService.SignIn(request)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "failed auth",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "success auth" + token,
	})
}
