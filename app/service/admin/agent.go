package admin

import (
	adminRepository "github.com/bagusyanuar/go-olin-bags/app/repositories/admin"
	"github.com/bagusyanuar/go-olin-bags/model"
)

type Agent interface {
	FindAll() ([]model.Agent, error)
}

type AgentService struct {
	AgentRepository adminRepository.Agent
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
