package admin

import (
	"fmt"
	"net/http"

	adminSvc "github.com/bagusyanuar/go-olin-bags/app/service/admin"
	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AgentController struct {
	AgentService adminSvc.Agent
}

func NewAgentController(agentService adminSvc.Agent) AgentController {
	return AgentController{AgentService: agentService}
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
