package agent

type PurchasingRequest struct {
	ProductionHouseID   string `json:"production_house_id" binding:"required,uuid4"`
	ShippingDestination string `json:"shipping_destination" binding:"required"`
	ShippingCost        uint64 `json:"shipping_cost" binding:"required,numeric"`
}
