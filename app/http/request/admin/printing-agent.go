package admin

type PrintingAgentRequest struct {
	ProductionHouseID string `json:"production_house_id" binding:"required,uuid4"`
	Name              string `json:"name" binding:"required"`
	Phone             string `json:"phone" binding:"required,e164"`
	Address           string `json:"address" binding:"required"`
}
