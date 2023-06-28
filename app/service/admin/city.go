package admin

import (
	request "github.com/bagusyanuar/go-olin-bags/app/http/request/admin"
	repository "github.com/bagusyanuar/go-olin-bags/app/repositories/admin"
	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/bagusyanuar/go-olin-bags/model"
	"github.com/google/uuid"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type City interface {
	FindAll(q string, limit, offset int) ([]model.City, error)
	FindByID(id string) (*model.City, error)
	Create(request request.CityRequest) (*model.City, error)
	Patch(id string, request request.CityRequest) error
	Delete(id string)
}

type CityService struct {
	CityRepository repository.City
}

// Create implements City.
func (svc *CityService) Create(request request.CityRequest) (*model.City, error) {
	provinceID, err := uuid.Parse(request.ProvinceID)
	if err != nil {
		return nil, err
	}
	e := model.City{
		ProvinceID: provinceID,
		Name:       cases.Title(language.Indonesian, cases.Compact).String(request.Name),
		Code:       request.Code,
	}
	return svc.CityRepository.Create(e)
}

// Delete implements City.
func (svc *CityService) Delete(id string) {
	panic("unimplemented")
}

// FindAll implements City.
func (svc *CityService) FindAll(q string, limit int, offset int) ([]model.City, error) {
	if limit == 0 {
		limit = common.DefaultLimit
	}
	return svc.CityRepository.FindAll(q, limit, offset)
}

// FindByID implements City.
func (svc *CityService) FindByID(id string) (*model.City, error) {
	return svc.CityRepository.FindByID(id)
}

// Patch implements City.
func (svc *CityService) Patch(id string, request request.CityRequest) error {
	panic("unimplemented")
}

func NewCityService(cityRepository repository.City) City {
	return &CityService{
		CityRepository: cityRepository,
	}
}
