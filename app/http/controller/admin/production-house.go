package admin

import (
	"net/http"
	"strconv"

	request "github.com/bagusyanuar/go-olin-bags/app/http/request/admin"
	service "github.com/bagusyanuar/go-olin-bags/app/service/admin"
	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/gin-gonic/gin"
)

type ProductionHouseController struct {
	ProductionHouseService service.ProductionHouse
	APIGroup               *gin.RouterGroup
}

func NewProductionHouseController(
	productionHouseService service.ProductionHouse,
	apiGroup *gin.RouterGroup,
) ProductionHouseController {
	return ProductionHouseController{
		ProductionHouseService: productionHouseService,
		APIGroup:               apiGroup,
	}
}

func (c *ProductionHouseController) RegisterRoutes() {
	route := c.APIGroup.Group("/production-house")
	{
		route.GET("/", c.FindAll)
		route.POST("/", c.Create)
		route.GET("/:id", c.FindByID)
	}
}

func (c *ProductionHouseController) FindAll(ctx *gin.Context) {
	q := ctx.Query("q")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	data, err := c.ProductionHouseService.FindAll(q, limit, offset)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, common.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: "internal server error",
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

func (c *ProductionHouseController) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")
	data, err := c.ProductionHouseService.FindByID(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, common.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: "internal server error",
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

func (c *ProductionHouseController) Create(ctx *gin.Context) {
	var req request.ProductionHouseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errorMessages := common.TranslateError(err, &req)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, common.APIResponse{
			Code:    http.StatusBadRequest,
			Message: "bad request",
			Data:    errorMessages,
		})
		return
	}

	data, err := c.ProductionHouseService.Create(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, common.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
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
