package admin

import (
	request "github.com/bagusyanuar/go-olin-bags/app/http/request/admin"
	repository "github.com/bagusyanuar/go-olin-bags/app/repositories/admin"
	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/bagusyanuar/go-olin-bags/model"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Province interface {
	FindAll(q string, limit, offset int) ([]model.Province, error)
	FindByID(id string) (*model.Province, error)
	Create(request request.ProvinceRequest) (*model.Province, error)
	Patch(id string, request request.ProvinceRequest) error
	Delete(id string)
}

type ProvinceService struct {
	ProvinceRepository repository.Province
}

// Create implements Province.
func (svc *ProvinceService) Create(request request.ProvinceRequest) (*model.Province, error) {
	e := model.Province{
		Name: cases.Title(language.Indonesian, cases.Compact).String(request.Name),
		Code: request.Code,
	}
	return svc.ProvinceRepository.Create(e)
}

// Delete implements Province.
func (svc *ProvinceService) Delete(id string) {
	panic("unimplemented")
}

// Patch implements Province.
func (svc *ProvinceService) Patch(id string, request request.ProvinceRequest) error {
	panic("unimplemented")
}

// FindAll implements Province.
func (svc *ProvinceService) FindAll(q string, limit, offset int) ([]model.Province, error) {
	if limit == 0 {
		limit = common.DefaultLimit
	}
	return svc.ProvinceRepository.FindAll(q, limit, offset)
}

// FindByID implements Province.
func (svc *ProvinceService) FindByID(id string) (*model.Province, error) {
	return svc.ProvinceRepository.FindByID(id)
}

func NewProvinceService(provinceRepository repository.Province) Province {
	return &ProvinceService{
		ProvinceRepository: provinceRepository,
	}
}
