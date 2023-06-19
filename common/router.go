package common

import "github.com/gin-gonic/gin"

type Route struct {
	Method  string
	Group   string
	Path    string
	Handler []gin.HandlerFunc
}
