package admin

type ItemRequest struct {
	MaterialID  string `json:"material_id" binding:"required,uuid4"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Price       int64  `json:"price" binding:"required"`
}
