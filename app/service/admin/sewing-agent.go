package admin

import (
	request "github.com/bagusyanuar/go-olin-bags/app/http/request/admin"
	repository "github.com/bagusyanuar/go-olin-bags/app/repositories/admin"
	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/bagusyanuar/go-olin-bags/model"
	"github.com/google/uuid"
)

type SewingAgent interface {
	FindAll(q string, limit, offset int) ([]model.SewingAgent, error)
	FindByID(id string) (*model.SewingAgent, error)
	Create(request request.SewingAgentRequest) (*model.SewingAgent, error)
}

type SewingAgentService struct {
	SewingAgentRepository repository.SewingAgent
}

// Create implements SewingAgent.
func (svc *SewingAgentService) Create(request request.SewingAgentRequest) (*model.SewingAgent, error) {
	productionHouseID, err := uuid.Parse(request.ProductionHouseID)
	if err != nil {
		return nil, err
	}

	e := model.SewingAgent{
		ProductionHouseID: productionHouseID,
		Name:              request.Name,
		Phone:             request.Phone,
		Address:           request.Address,
	}
	return svc.SewingAgentRepository.Create(e)
}

// FindAll implements SewingAgent.
func (svc *SewingAgentService) FindAll(q string, limit int, offset int) ([]model.SewingAgent, error) {
	if limit == 0 {
		limit = common.DefaultLimit
	}
	return svc.SewingAgentRepository.FindAll(q, limit, offset)
}

// FindByID implements SewingAgent.
func (svc *SewingAgentService) FindByID(id string) (*model.SewingAgent, error) {
	return svc.SewingAgentRepository.FindByID(id)
}

func NewSewingAgentService(sewingAgentRepository repository.SewingAgent) SewingAgent {
	return &SewingAgentService{
		SewingAgentRepository: sewingAgentRepository,
	}
}
