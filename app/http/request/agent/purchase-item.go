package agent

type PurchaseItemRequest struct {
	ItemID string `json:"item_id" binding:"required,uuid4"`
	Qty    uint32 `json:"qty" binding:"required,numeric"`
}
