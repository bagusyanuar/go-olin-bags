package admin

type CityRequest struct {
	ProvinceID string `json:"province_id" binding:"required,uuid4"`
	Name       string `json:"name" binding:"required"`
	Code       string `json:"code" binding:"required,numeric,len=6"`
}
