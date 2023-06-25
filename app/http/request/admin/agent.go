package admin

type AgentRequest struct {
	Email     string  `json:"email" binding:"required,email"`
	Username  string  `json:"username" binding:"required"`
	Password  string  `json:"password" binding:"required"`
	CityID    string  `json:"city_id" binding:"required,uuid4"`
	Name      string  `json:"name" binding:"required"`
	Phone     string  `json:"phone" binding:"required,e164"`
	Address   string  `json:"address" binding:"required"`
	Latitude  float64 `json:"latitude" binding:"required,latitude"`
	Longitude float64 `json:"longitude" binding:"required,longitude"`
}
