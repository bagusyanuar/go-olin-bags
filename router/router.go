package router

import (
	"net/http"

	"github.com/bagusyanuar/go-olin-bags/app/http/controller"
	"github.com/bagusyanuar/go-olin-bags/common"
)

func Routers(
	dashboardController *controller.DashboardController,
) []*common.Route {
	return []*common.Route{
		{
			Method:  http.MethodGet,
			Path:    "/",
			Handler: dashboardController.Index,
		},
	}
}
