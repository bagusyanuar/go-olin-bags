package productionhouse

import (
	"errors"

	request "github.com/bagusyanuar/go-olin-bags/app/http/request/production-house"
	repository "github.com/bagusyanuar/go-olin-bags/app/repositories/production-house"
	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/bagusyanuar/go-olin-bags/model"
	"gorm.io/gorm"
)

type PrintingAgent interface {
	FindAll(authorizedID, q string, limit, offset int) ([]model.PrintingAgent, error)
	FindByID(authorizedID, id string) (*model.PrintingAgent, error)
	Create(authorizedID string, request request.PrintingAgentRequest) (*model.PrintingAgent, error)
}

type PrintingAgentService struct {
	PrintingAgentRepository repository.PrintingAgent
	ProfileRepository       repository.Profile
}

// Create implements PrintingAgent.
func (svc *PrintingAgentService) Create(authorizedID string, request request.PrintingAgentRequest) (*model.PrintingAgent, error) {
	productionHouse, err := svc.ProfileRepository.GetProfile(authorizedID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("production house not found")
		}
		return nil, err
	}
	entity := model.PrintingAgent{
		ProductionHouseID: productionHouse.ID,
		Name:              request.Name,
		Phone:             request.Phone,
		Address:           request.Address,
	}
	return svc.PrintingAgentRepository.Create(entity)
}

// FindAll implements PrintingAgent.
func (svc *PrintingAgentService) FindAll(authorizedID string, q string, limit int, offset int) ([]model.PrintingAgent, error) {
	productionHouse, err := svc.ProfileRepository.GetProfile(authorizedID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("production house not found")
		}
		return nil, err
	}
	if limit == 0 {
		limit = common.DefaultLimit
	}
	return svc.PrintingAgentRepository.FindAll(productionHouse.ID.String(), q, limit, offset)
}

// FindByID implements PrintingAgent.
func (svc *PrintingAgentService) FindByID(authorizedID string, id string) (*model.PrintingAgent, error) {
	productionHouse, err := svc.ProfileRepository.GetProfile(authorizedID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("production house not found")
		}
		return nil, err
	}
	return svc.PrintingAgentRepository.FindByID(productionHouse.ID.String(), id)
}

func NewPrintingAgentService(
	printingAgentRepository repository.PrintingAgent,
	profileRepository repository.Profile,
) PrintingAgent {
	return &PrintingAgentService{
		PrintingAgentRepository: printingAgentRepository,
		ProfileRepository:       profileRepository,
	}
}
