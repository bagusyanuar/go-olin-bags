package admin

import (
	"fmt"
	"net/http"

	adminSvc "github.com/bagusyanuar/go-olin-bags/app/service/admin"
	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/gin-gonic/gin"
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
		ctx.JSON(500, common.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("internal server error (%s)", err.Error()),
			Data:    nil,
		})
		return
	}
	ctx.JSON(200, common.APIResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}
