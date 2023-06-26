package admin

import (
	request "github.com/bagusyanuar/go-olin-bags/app/http/request/admin"
	repository "github.com/bagusyanuar/go-olin-bags/app/repositories/admin"
	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/bagusyanuar/go-olin-bags/model"
	"github.com/google/uuid"
)

type PrintingAgent interface {
	FindAll(q string, limit, offset int) ([]model.PrintingAgent, error)
	FindByID(id string) (*model.PrintingAgent, error)
	Create(request request.PrintingAgentRequest) (*model.PrintingAgent, error)
}

type PrintingAgentService struct {
	PrintingAgentRepository repository.PrintingAgent
}

// Create implements PrintingAgent.
func (svc *PrintingAgentService) Create(request request.PrintingAgentRequest) (*model.PrintingAgent, error) {
	productionHouseID, err := uuid.Parse(request.ProductionHouseID)
	if err != nil {
		return nil, err
	}

	e := model.PrintingAgent{
		ProductionHouseID: productionHouseID,
		Name:              request.Name,
		Phone:             request.Phone,
		Address:           request.Address,
	}
	return svc.PrintingAgentRepository.Create(e)
}

// FindAll implements PrintingAgent.
func (svc *PrintingAgentService) FindAll(q string, limit int, offset int) ([]model.PrintingAgent, error) {
	if limit == 0 {
		limit = common.DefaultLimit
	}
	return svc.PrintingAgentRepository.FindAll(q, limit, offset)
}

// FindByID implements PrintingAgent.
func (svc *PrintingAgentService) FindByID(id string) (*model.PrintingAgent, error) {
	return svc.PrintingAgentRepository.FindByID(id)
}

func NewPrintingAgentService(printingAgentRepository repository.PrintingAgent) PrintingAgent {
	return &PrintingAgentService{
		PrintingAgentRepository: printingAgentRepository,
	}
}
