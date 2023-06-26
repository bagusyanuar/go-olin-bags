package admin

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bagusyanuar/go-olin-bags/app/http/middleware"
	request "github.com/bagusyanuar/go-olin-bags/app/http/request/admin"
	service "github.com/bagusyanuar/go-olin-bags/app/service/admin"
	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ItemController struct {
	ItemService service.Item
	APIGroup    *gin.RouterGroup
	Middleware  middleware.AuthMiddleware
}

func NewItemController(
	itemService service.Item,
	apiGroup *gin.RouterGroup,
	middleware middleware.AuthMiddleware,
) ItemController {
	return ItemController{
		ItemService: itemService,
		APIGroup:    apiGroup,
		Middleware:  middleware,
	}
}

func (c *ItemController) RegisterRoutes() {
	route := c.APIGroup.Group("/item")
	{
		route.GET("/", c.Middleware.IsAuth(), c.FindAll)
		route.POST("/", c.Middleware.IsAuth(), c.Create)
		route.GET("/:id", c.Middleware.IsAuth(), c.FindByID)
	}
}

func (c *ItemController) FindAll(ctx *gin.Context) {
	q := ctx.Query("q")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	data, err := c.ItemService.FindAll(q, limit, offset)
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

func (c *ItemController) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")
	data, err := c.ItemService.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, common.APIResponse{
				Code:    http.StatusNotFound,
				Message: "data not found",
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
	ctx.JSON(http.StatusOK, common.APIResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}

func (c *ItemController) Create(ctx *gin.Context) {
	var req request.ItemRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errorMessages := common.TranslateError(err, &req)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, common.APIResponse{
			Code:    http.StatusBadRequest,
			Message: "bad request",
			Data:    errorMessages,
		})
		return
	}

	data, err := c.ItemService.Create(req)
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
