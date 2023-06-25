package productionhouse

import (
	"errors"

	request "github.com/bagusyanuar/go-olin-bags/app/http/request/production-house"
	repository "github.com/bagusyanuar/go-olin-bags/app/repositories/production-house"
	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/bagusyanuar/go-olin-bags/model"
	"gorm.io/gorm"
)

type SewingAgent interface {
	FindAll(authorizedID, q string, limit, offset int) ([]model.SewingAgent, error)
	FindByID(authorizedID, id string) (*model.SewingAgent, error)
	Create(authorizedID string, request request.SewingAgentRequest) (*model.SewingAgent, error)
}

type SewingAgentService struct {
	SewingAgentRepository repository.SewingAgent
	ProfileRepository     repository.Profile
}

// Create implements SewingAgent.
func (svc *SewingAgentService) Create(authorizedID string, request request.SewingAgentRequest) (*model.SewingAgent, error) {
	productionHouse, err := svc.ProfileRepository.GetProfile(authorizedID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("production house not found")
		}
		return nil, err
	}
	entity := model.SewingAgent{
		ProductionHouseID: productionHouse.ID,
		Name:              request.Name,
		Phone:             request.Phone,
		Address:           request.Address,
	}
	return svc.SewingAgentRepository.Create(entity)
}

// FindAll implements SewingAgent.
func (svc *SewingAgentService) FindAll(authorizedID string, q string, limit int, offset int) ([]model.SewingAgent, error) {
	if limit == 0 {
		limit = common.DefaultLimit
	}
	return svc.SewingAgentRepository.FindAll(authorizedID, q, limit, offset)
}

// FindByID implements SewingAgent.
func (svc *SewingAgentService) FindByID(authorizedID string, id string) (*model.SewingAgent, error) {
	return svc.SewingAgentRepository.FindByID(authorizedID, id)
}

func NewSewingAgentService(
	sewingAgentRepository repository.SewingAgent,
	profileRepository repository.Profile,
) SewingAgent {
	return &SewingAgentService{
		SewingAgentRepository: sewingAgentRepository,
		ProfileRepository:     profileRepository,
	}
}
