package admin

import (
	request "github.com/bagusyanuar/go-olin-bags/app/http/request/admin"
	repository "github.com/bagusyanuar/go-olin-bags/app/repositories/admin"
	"github.com/bagusyanuar/go-olin-bags/model"
)

type Province interface {
	FindAll() ([]model.Province, error)
	FindByID(id string) (*model.Province, error)
	Create(request request.ProvinceRequest) error
	Patch(id string, request request.ProvinceRequest) error
	Delete(id string)
}

type ProvinceService struct {
	ProvinceRepository repository.Province
}

// Create implements Province.
func (svc *ProvinceService) Create(request request.ProvinceRequest) error {
	panic("unimplemented")
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
func (svc *ProvinceService) FindAll() ([]model.Province, error) {
	panic("unimplemented")
}

// FindByID implements Province.
func (svc *ProvinceService) FindByID(id string) (*model.Province, error) {
	panic("unimplemented")
}

func NewProvinceService(provinceRepository repository.Province) Province {
	return &ProvinceService{
		ProvinceRepository: provinceRepository,
	}
}
