package productionhouse

type PrintingAgentRequest struct {
	Name    string `json:"name" binding:"required"`
	Phone   string `json:"phone" binding:"required,e164"`
	Address string `json:"address" binding:"required"`
}
