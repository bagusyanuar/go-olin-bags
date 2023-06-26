package admin

type MaterialRequest struct {
	Name string `json:"name" binding:"required"`
}
