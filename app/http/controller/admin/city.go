package admin

import (
	"net/http"
	"strconv"

	request "github.com/bagusyanuar/go-olin-bags/app/http/request/admin"
	service "github.com/bagusyanuar/go-olin-bags/app/service/admin"
	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/gin-gonic/gin"
)

type CityController struct {
	CityService service.City
	APIGroup    *gin.RouterGroup
}

func NewCityController(
	cityService service.City,
	apiGroup *gin.RouterGroup,
) CityController {
	return CityController{
		CityService: cityService,
		APIGroup:    apiGroup,
	}
}

func (c *CityController) RegisterRoutes() {
	route := c.APIGroup.Group("/city")
	{
		route.GET("/", c.FindAll)
		route.POST("/", c.Create)
		route.GET("/:id", c.FindByID)
	}
}

func (c *CityController) FindAll(ctx *gin.Context) {
	q := ctx.Query("q")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	data, err := c.CityService.FindAll(q, limit, offset)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, common.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: "bad request",
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

func (c *CityController) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")
	data, err := c.CityService.FindByID(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, common.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: "bad request",
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

func (c *CityController) Create(ctx *gin.Context) {
	var req request.CityRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errorMessages := common.TranslateError(err, &req)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, common.APIResponse{
			Code:    http.StatusBadRequest,
			Message: "bad request",
			Data:    errorMessages,
		})
		return
	}

	err := c.CityService.Create(req)
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
		Data:    nil,
	})
}
