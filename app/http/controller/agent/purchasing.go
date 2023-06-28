package agent

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bagusyanuar/go-olin-bags/app/http/middleware"
	request "github.com/bagusyanuar/go-olin-bags/app/http/request/agent"
	service "github.com/bagusyanuar/go-olin-bags/app/service/agent"
	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/gin-gonic/gin"
)

type PurchasingController struct {
	PurchasingService service.Purchasing
	APIGroup          *gin.RouterGroup
	Middleware        middleware.AuthMiddleware
}

func NewPurchasingController(
	purchasingService service.Purchasing,
	apiGroup *gin.RouterGroup,
	middleware middleware.AuthMiddleware,
) PurchasingController {
	return PurchasingController{
		PurchasingService: purchasingService,
		APIGroup:          apiGroup,
		Middleware:        middleware,
	}
}

func (c *PurchasingController) RegisterRoutes() {
	route := c.APIGroup.Group("/purchasing")
	{
		route.GET("/", c.Middleware.Auth, c.List)
		route.POST("/", c.Middleware.Auth, c.Checkout)
	}
}

func (c *PurchasingController) List(ctx *gin.Context) {
	accountID := common.GetAuthorizedID(ctx)
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	data, err := c.PurchasingService.List(accountID, limit, offset)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, common.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("internal server error (%s)", err.Error()),
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, common.APIResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}

func (c *PurchasingController) Checkout(ctx *gin.Context) {
	accountID := common.GetAuthorizedID(ctx)
	var req request.PurchasingRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errorMessages := common.TranslateError(err, &req)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, common.APIResponse{
			Code:    http.StatusBadRequest,
			Message: "bad request",
			Data:    errorMessages,
		})
		return
	}

	data, err := c.PurchasingService.Checkout(accountID, req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, common.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("internal server error (%s)", err.Error()),
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusCreated, common.APIResponse{
		Code:    http.StatusCreated,
		Message: "success",
		Data:    data,
	})
}
