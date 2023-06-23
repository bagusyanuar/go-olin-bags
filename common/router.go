package common

import (
	"reflect"

	"github.com/gin-gonic/gin"
)

const (
	RegisterRoutesMethod = "RegisterRoutes"
)

type Route struct {
	Method      string
	Group       string
	Path        string
	Middlewares []gin.HandlerFunc
	Handler     gin.HandlerFunc
}

func RegisterRoutes(controllers ...any) {
	for _, controller := range controllers {
		refType := reflect.TypeOf(controller)
		_, ok := refType.MethodByName(RegisterRoutesMethod)
		if ok {
			reflect.ValueOf(controller).MethodByName(RegisterRoutesMethod).Call([]reflect.Value{})
		}
	}
}
