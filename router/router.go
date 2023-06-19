package router

import (
	"net/http"

	"github.com/bagusyanuar/go-olin-bags/app/http/controller"
	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/gin-gonic/gin"
)

const (
	AuthPath  = "/auth"
	AdminPath = "/admin"
)

func Routers(
	welcomeController *controller.WelcomeController,
	AuthController *controller.AuthController,
) []*common.Route {
	return []*common.Route{
		{
			Method: http.MethodGet,
			Path:   "/",
			Handler: []gin.HandlerFunc{
				welcomeController.Index,
			},
		},
		{
			Method: http.MethodPost,
			Group:  AuthPath,
			Path:   "/sign-in",
			Handler: []gin.HandlerFunc{
				AuthController.SignIn,
			},
		},
	}
}
