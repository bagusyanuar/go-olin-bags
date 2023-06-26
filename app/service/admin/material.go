package admin

import (
	request "github.com/bagusyanuar/go-olin-bags/app/http/request/admin"
	repository "github.com/bagusyanuar/go-olin-bags/app/repositories/admin"
	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/bagusyanuar/go-olin-bags/model"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Material interface {
	FindAll(q string, limit, offset int) ([]model.Material, error)
	FindByID(id string) (*model.Material, error)
	Create(request request.MaterialRequest) (*model.Material, error)
}

type MaterialService struct {
	MaterialRepository repository.Material
}

// Create implements Material.
func (svc *MaterialService) Create(request request.MaterialRequest) (*model.Material, error) {
	entity := model.Material{
		Name: cases.Title(language.Indonesian, cases.Compact).String(request.Name),
	}
	return svc.MaterialRepository.Create(entity)
}

// FindAll implements Material.
func (svc *MaterialService) FindAll(q string, limit int, offset int) ([]model.Material, error) {
	if limit == 0 {
		limit = common.DefaultLimit
	}
	return svc.MaterialRepository.FindAll(q, limit, offset)
}

// FindByID implements Material.
func (svc *MaterialService) FindByID(id string) (*model.Material, error) {
	return svc.MaterialRepository.FindByID(id)
}

func NewMaterialService(materialRepository repository.Material) Material {
	return &MaterialService{
		MaterialRepository: materialRepository,
	}
}
