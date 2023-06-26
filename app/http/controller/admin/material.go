package admin

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bagusyanuar/go-olin-bags/app/http/middleware"
	request "github.com/bagusyanuar/go-olin-bags/app/http/request/admin"
	service "github.com/bagusyanuar/go-olin-bags/app/service/admin"
	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/gin-gonic/gin"
)

type MaterialController struct {
	MaterialService service.Material
	APIGroup        *gin.RouterGroup
	Middleware      middleware.AuthMiddleware
}

func NewMaterialController(
	materialService service.Material,
	apiGroup *gin.RouterGroup,
	middleware middleware.AuthMiddleware,
) MaterialController {
	return MaterialController{
		MaterialService: materialService,
		APIGroup:        apiGroup,
		Middleware:      middleware,
	}
}

func (c *MaterialController) RegisterRoutes() {
	route := c.APIGroup.Group("/material")
	{
		route.GET("/", c.Middleware.IsAuth(), c.FindAll)
		route.POST("/", c.Middleware.IsAuth(), c.Create)
		route.GET("/:id", c.Middleware.IsAuth(), c.FindByID)
	}
}

func (c *MaterialController) FindAll(ctx *gin.Context) {
	q := ctx.Query("q")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	data, err := c.MaterialService.FindAll(q, limit, offset)
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

func (c *MaterialController) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")
	data, err := c.MaterialService.FindByID(id)
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

func (c *MaterialController) Create(ctx *gin.Context) {
	var req request.MaterialRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errorMessages := common.TranslateError(err, &req)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, common.APIResponse{
			Code:    http.StatusBadRequest,
			Message: "bad request",
			Data:    errorMessages,
		})
		return
	}

	data, err := c.MaterialService.Create(req)
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
