package admin

import (
	"fmt"
	"net/http"

	request "github.com/bagusyanuar/go-olin-bags/app/http/request/admin"
	service "github.com/bagusyanuar/go-olin-bags/app/service/admin"
	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AgentController struct {
	AgentService service.Agent
	APIGroup     *gin.RouterGroup
}

func NewAgentController(
	agentService service.Agent,
	apiGroup *gin.RouterGroup,
) AgentController {
	return AgentController{
		AgentService: agentService,
		APIGroup:     apiGroup,
	}
}

func (c *AgentController) RegisterRoutes() {
	route := c.APIGroup.Group("/agent")
	{
		route.GET("/", c.FindAll)
		route.POST("/", c.Create)
		route.GET("/:id", c.FindByID)
	}
}

func (c *AgentController) FindAll(ctx *gin.Context) {
	data, err := c.AgentService.FindAll()
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

func (c *AgentController) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")
	data, err := c.AgentService.FindByID(id)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ctx.AbortWithStatusJSON(http.StatusNotFound, common.APIResponse{
				Code:    http.StatusNotFound,
				Message: "data not found",
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
		Data:    data,
	})
}

func (c *AgentController) Create(ctx *gin.Context) {
	var req request.AgentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errorMessages := common.TranslateError(err, &req)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, common.APIResponse{
			Code:    http.StatusBadRequest,
			Message: "bad request",
			Data:    errorMessages,
		})
		return
	}

	data, err := c.AgentService.Create(req)
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
