package router

import (
	"net/http"

	"github.com/bagusyanuar/go-olin-bags/app/http/controller"
	"github.com/bagusyanuar/go-olin-bags/common"
)

const (
	AuthPath  = "/auth"
	AdminPath = "/admin"
)

func PublicRoutes(
	welcomeController *controller.WelcomeController,
	AuthController *controller.AuthController,
) []*common.Route {
	return []*common.Route{
		{Method: http.MethodGet, Path: "/", Handler: welcomeController.Index},
		//route auth
		{Method: http.MethodPost, Group: AuthPath, Path: "/sign-in", Handler: AuthController.SignIn},
	}
}

// func AdminRouter(
// 	agentController *adminCtrl.AgentController,
// ) []*common.Route {
// 	agentGroup := fmt.Sprintf("%s%s", AdminPath, AgentPath)
// 	return []*common.Route{

// 		{Method: http.MethodGet, Group: agentGroup, Path: "/", Handler: []gin.HandlerFunc{agentController.FindAll}},
// 	}
// }
