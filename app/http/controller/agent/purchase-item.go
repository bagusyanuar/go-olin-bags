package agent

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/bagusyanuar/go-olin-bags/app/http/middleware"
	request "github.com/bagusyanuar/go-olin-bags/app/http/request/agent"
	service "github.com/bagusyanuar/go-olin-bags/app/service/agent"
	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PurchaseItemController struct {
	PurchaseItemService service.PurchaseItem
	APIGroup            *gin.RouterGroup
	Middleware          middleware.AuthMiddleware
}

func NewPurchaseItemController(
	purchaseItemService service.PurchaseItem,
	apiGroup *gin.RouterGroup,
	middleware middleware.AuthMiddleware,
) PurchaseItemController {
	return PurchaseItemController{
		PurchaseItemService: purchaseItemService,
		APIGroup:            apiGroup,
		Middleware:          middleware,
	}
}

func (c *PurchaseItemController) RegisterRoutes() {
	route := c.APIGroup.Group("/purchase-item")
	{
		route.GET("/", c.Middleware.Auth, c.List)
		route.POST("/", c.Middleware.Auth, c.AddItem)
	}
}

func (c *PurchaseItemController) List(ctx *gin.Context) {
	accountID := common.GetAuthorizedID(ctx)
	data, err := c.PurchaseItemService.List(accountID)
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

func (c *PurchaseItemController) AddItem(ctx *gin.Context) {
	accountID := common.GetAuthorizedID(ctx)
	var req request.PurchaseItemRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errorMessages := common.TranslateError(err, &req)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, common.APIResponse{
			Code:    http.StatusBadRequest,
			Message: "bad request",
			Data:    errorMessages,
		})
		return
	}

	data, err := c.PurchaseItemService.AddItem(accountID, req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, common.APIResponse{
				Code:    http.StatusNotFound,
				Message: "item not found",
				Data:    nil,
			})
			return
		}
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
