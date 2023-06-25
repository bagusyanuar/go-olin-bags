package common

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var ErrorPasswordNotMatch = errors.New("password did not match")
var ErrorUnauthorized = errors.New("unauthorized")

func Catch(ctx *gin.Context) {
	if r := recover(); r != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, APIResponse{
			Code:    http.StatusInternalServerError,
			Message: "internal server error",
			Data:    nil,
		})
		return
	}
}
