package admin

type ProvinceRequest struct {
	Name string `json:"name" binding:"required"`
	Code string `json:"code" binding:"required,numeric,len=4"`
}
