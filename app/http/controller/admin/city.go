package admin

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type CityController struct{}

func (c *CityController) RegisterRoutes() {
	fmt.Println("Method Register Routes")
}

func (c *CityController) FindAll(ctx *gin.Context) {

}

func NewCityController() CityController {
	return CityController{}
}
