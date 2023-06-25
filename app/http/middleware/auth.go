package middleware

import (
	"net/http"

	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/bagusyanuar/go-olin-bags/config"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	Config *config.JWT
}

func NewAuthMiddleware(config *config.JWT) AuthMiddleware {
	return AuthMiddleware{Config: config}
}

func (m *AuthMiddleware) IsAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.Request.Header.Get("Authorization")
		if authorization == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, common.APIResponse{
				Code:    http.StatusUnauthorized,
				Message: "unauthorized",
			})
			return
		}
		bearer := string(authorization[0:7])
		token := string(authorization[7:])
		if bearer != "Bearer " {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, common.APIResponse{
				Code:    http.StatusUnauthorized,
				Message: "error bearer type",
			})
			return
		}
		jwtClaim, err := common.VerifyToken(token, m.Config)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, common.APIResponse{
				Code:    http.StatusUnauthorized,
				Data:    nil,
				Message: err.Error(),
			})
			return
		}
		ctx.Set("user", jwtClaim)
		ctx.Next()
	}
}
