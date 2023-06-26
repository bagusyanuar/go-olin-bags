package admin

import (
	request "github.com/bagusyanuar/go-olin-bags/app/http/request/admin"
	repository "github.com/bagusyanuar/go-olin-bags/app/repositories/admin"
	"github.com/bagusyanuar/go-olin-bags/model"
	"github.com/google/uuid"
)

type ProductionHouse interface {
	FindAll(q string, limit, offset int) ([]model.ProductionHouse, error)
	FindByID(id string) (*model.ProductionHouse, error)
	Create(request request.ProductionHouseRequest) (*model.ProductionHouse, error)
	Patch(id string, request request.ProductionHouseRequest) error
	Delete(id string)
}

type ProductionHouseService struct {
	ProductionHouseRepository repository.ProductionHouse
}

// Create implements ProductionHouse.
func (svc *ProductionHouseService) Create(request request.ProductionHouseRequest) (*model.ProductionHouse, error) {
	cityID, err := uuid.Parse(request.CityID)
	if err != nil {
		return nil, err
	}

	e := model.ProductionHouse{
		CityID:    cityID,
		Name:      request.Name,
		Phone:     request.Phone,
		Address:   request.Address,
		Latitude:  request.Latitude,
		Longitude: request.Longitude,
	}
	return svc.ProductionHouseRepository.Create(e)
}

// Delete implements ProductionHouse.
func (svc *ProductionHouseService) Delete(id string) {
	panic("unimplemented")
}

// FindAll implements ProductionHouse.
func (svc *ProductionHouseService) FindAll(q string, limit int, offset int) ([]model.ProductionHouse, error) {
	panic("unimplemented")
}

// FindByID implements ProductionHouse.
func (svc *ProductionHouseService) FindByID(id string) (*model.ProductionHouse, error) {
	panic("unimplemented")
}

// Patch implements ProductionHouse.
func (svc *ProductionHouseService) Patch(id string, request request.ProductionHouseRequest) error {
	panic("unimplemented")
}

func NewProductionHouseService(productionHouseRepository repository.ProductionHouse) ProductionHouse {
	return &ProductionHouseService{
		ProductionHouseRepository: productionHouseRepository,
	}
}
