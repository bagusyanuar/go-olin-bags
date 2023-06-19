package router

import (
	"fmt"
	"net/http"

	"github.com/bagusyanuar/go-olin-bags/app/http/controller"
	adminCtrl "github.com/bagusyanuar/go-olin-bags/app/http/controller/admin"
	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/gin-gonic/gin"
)

const (
	AuthPath  = "/auth"
	AdminPath = "/admin"
	AgentPath = "/agent"
)

func Routers(
	welcomeController *controller.WelcomeController,
	AuthController *controller.AuthController,
	agentAdminController *adminCtrl.AgentController,

) []*common.Route {
	return []*common.Route{
		{Method: http.MethodGet, Path: "/", Handler: []gin.HandlerFunc{welcomeController.Index}},
		//route auth
		{Method: http.MethodPost, Group: AuthPath, Path: "/sign-in", Handler: []gin.HandlerFunc{AuthController.SignIn}},
		//route admin agent
		{Method: http.MethodGet, Group: fmt.Sprintf("%s%s", AdminPath, AgentPath), Path: "/", Handler: []gin.HandlerFunc{agentAdminController.FindAll}},
	}
}
