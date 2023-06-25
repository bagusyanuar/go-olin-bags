package productionhouse

import (
	"fmt"
	"net/http"

	request "github.com/bagusyanuar/go-olin-bags/app/http/request/production-house"
	service "github.com/bagusyanuar/go-olin-bags/app/service/production-house"
	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	AuthService service.Auth
	APIGroup    *gin.RouterGroup
}

func NewAuthController(
	authService service.Auth,
	apiGroup *gin.RouterGroup,
) AuthController {
	return AuthController{
		AuthService: authService,
		APIGroup:    apiGroup,
	}
}

func (c *AuthController) RegisterRoutes() {
	route := c.APIGroup.Group("/auth")
	{
		route.POST("/sign-in", c.SignIn)
	}
}

func (c *AuthController) SignIn(ctx *gin.Context) {
	var req request.SignInRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errorMessages := common.TranslateError(err, &req)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, common.APIResponse{
			Code:    http.StatusBadRequest,
			Message: "bad request",
			Data:    errorMessages,
		})
		return
	}

	token, err := c.AuthService.SignIn(req)
	if err != nil {
		switch err {
		case common.ErrorPasswordNotMatch:
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, common.APIResponse{
				Code:    http.StatusUnauthorized,
				Message: err.Error(),
				Data:    nil,
			})
			return
		case gorm.ErrRecordNotFound:
			ctx.AbortWithStatusJSON(http.StatusNotFound, common.APIResponse{
				Code:    http.StatusNotFound,
				Message: "user not found",
				Data:    nil,
			})
			return
		default:
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, common.APIResponse{
				Code:    http.StatusInternalServerError,
				Message: fmt.Sprintf("internal server error (%s)", err.Error()),
				Data:    nil,
			})
			return
		}
	}
	ctx.JSON(200, common.APIResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    token,
	})
}
