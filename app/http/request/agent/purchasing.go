package agent

type PurchasingRequest struct {
	ShippingDestination string `json:"shipping_destination" binding:"required"`
	ShippingCost        uint64 `json:"shipping_cost" binding:"required,numeric"`
}
