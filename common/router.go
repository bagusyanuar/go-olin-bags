package common

import "github.com/gin-gonic/gin"

type Route struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

func BuildRoute(route *gin.Engine, routers []*Route) {
	for _, r := range routers {
		route.Handle(r.Method, r.Path, r.Handler)
	}
}
