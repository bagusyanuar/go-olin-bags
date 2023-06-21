package admin

import (
	repository "github.com/bagusyanuar/go-olin-bags/app/repositories/admin"
	"github.com/bagusyanuar/go-olin-bags/model"
)

type Province interface {
	FindAll() ([]model.Province, error)
	FindByID(id string) (*model.Province, error)
}

type ProvinceService struct {
	ProvinceRepository repository.Province
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
