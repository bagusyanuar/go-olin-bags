package admin

import (
	"encoding/json"
	"errors"

	request "github.com/bagusyanuar/go-olin-bags/app/http/request/admin"
	adminRepository "github.com/bagusyanuar/go-olin-bags/app/repositories/admin"
	"github.com/bagusyanuar/go-olin-bags/model"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Agent interface {
	FindAll() ([]model.Agent, error)
	FindByID(id string) (*model.Agent, error)
	Create(request request.AgentRequest) (*model.Agent, error)
}

type AgentService struct {
	AgentRepository adminRepository.Agent
}

// Create implements Agent.
func (svc *AgentService) Create(request request.AgentRequest) (*model.Agent, error) {
	if request.Password == "" {
		return nil, errors.New("password cannot empty")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), 13)
	if err != nil {
		return nil, err
	}
	password := string(hash)

	roles, _ := json.Marshal([]string{"agent"})
	user := model.User{
		Email:    request.Email,
		Username: request.Username,
		Password: &password,
		Roles:    roles,
	}

	cityID, err := uuid.Parse(request.CityID)
	if err != nil {
		return nil, err
	}

	e := model.Agent{
		CityID:    cityID,
		Name:      request.Name,
		Phone:     request.Phone,
		Address:   request.Address,
		Latitude:  request.Latitude,
		Longitude: request.Longitude,
		User:      &user,
	}
	return svc.AgentRepository.Create(e)
}

// FindByID implements Agent.
func (svc *AgentService) FindByID(id string) (*model.Agent, error) {
	return svc.AgentRepository.FindByID(id)
}

// FindAll implements Agent.
func (svc *AgentService) FindAll() ([]model.Agent, error) {
	return svc.AgentRepository.FindAll()
}

func NewAgentService(agentRepository adminRepository.Agent) Agent {
	return &AgentService{
		AgentRepository: agentRepository,
	}
}
