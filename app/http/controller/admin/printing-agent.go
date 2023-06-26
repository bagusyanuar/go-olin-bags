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

type PrintingAgentController struct {
	PrintingAgentService service.PrintingAgent
	APIGroup             *gin.RouterGroup
	Middleware           middleware.AuthMiddleware
}

func NewPrintingAgentController(
	printingAgentService service.PrintingAgent,
	apiGroup *gin.RouterGroup,
	middleware middleware.AuthMiddleware,
) PrintingAgentController {
	return PrintingAgentController{
		PrintingAgentService: printingAgentService,
		APIGroup:             apiGroup,
		Middleware:           middleware,
	}
}

func (c *PrintingAgentController) RegisterRoutes() {
	route := c.APIGroup.Group("/printing-agent")
	{
		route.GET("/", c.Middleware.Auth, c.FindAll)
		route.POST("/", c.Middleware.Auth, c.Create)
		route.GET("/:id", c.Middleware.Auth, c.FindByID)
	}
}

func (c *PrintingAgentController) FindAll(ctx *gin.Context) {
	q := ctx.Query("q")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	data, err := c.PrintingAgentService.FindAll(q, limit, offset)
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

func (c *PrintingAgentController) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")
	data, err := c.PrintingAgentService.FindByID(id)
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

func (c *PrintingAgentController) Create(ctx *gin.Context) {
	var req request.PrintingAgentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errorMessages := common.TranslateError(err, &req)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, common.APIResponse{
			Code:    http.StatusBadRequest,
			Message: "bad request",
			Data:    errorMessages,
		})
		return
	}

	data, err := c.PrintingAgentService.Create(req)
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
