package router

import (
	"net/http"

	controller "github.com/bagusyanuar/go-olin-bags/app/http/controller/admin"
	"github.com/bagusyanuar/go-olin-bags/common"
)

const (
	AgentPath    = "/agent"
	ProvincePath = "/province"
)

func AdminRoutes(
	agentController *controller.AgentController,
	provinceController *controller.ProvinceController,
) []*common.Route {
	return []*common.Route{
		{
			Method:  http.MethodGet,
			Group:   AgentPath,
			Path:    "/",
			Handler: agentController.FindAll,
		},
		{
			Method:  http.MethodGet,
			Group:   AgentPath,
			Path:    "/:id",
			Handler: agentController.FindByID,
		},
		{
			Method:  http.MethodPost,
			Group:   ProvincePath,
			Path:    "/",
			Handler: provinceController.Create,
		},
	}
}
