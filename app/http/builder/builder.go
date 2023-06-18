package builder

import (
	"github.com/bagusyanuar/go-olin-bags/app/http/controller"
	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/bagusyanuar/go-olin-bags/router"
)

func BuildRoute() []*common.Route {
	dashboardController := controller.NewDashboardController()
	return router.Routers(
		&dashboardController,
	)
}
