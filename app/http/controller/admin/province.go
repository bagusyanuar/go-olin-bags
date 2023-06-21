package admin

import (
	"net/http"

	request "github.com/bagusyanuar/go-olin-bags/app/http/request/admin"
	service "github.com/bagusyanuar/go-olin-bags/app/service/admin"
	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/gin-gonic/gin"
)

type ProvinceController struct {
	ProvinceService service.Province
}

func NewProvinceController(provinceService service.Province) ProvinceController {
	return ProvinceController{ProvinceService: provinceService}
}

func (c *ProvinceController) Create(ctx *gin.Context) {
	var req request.ProvinceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errorMessages := common.TranslateError(err, &req)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, common.APIResponse{
			Code:    http.StatusBadRequest,
			Message: "bad request",
			Data:    errorMessages,
		})
		return
	}
	ctx.JSON(http.StatusCreated, common.APIResponse{
		Code:    http.StatusCreated,
		Message: "success",
		Data:    nil,
	})
}
