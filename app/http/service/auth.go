package service

import (
	"github.com/bagusyanuar/go-olin-bags/app/http/request"
	"github.com/bagusyanuar/go-olin-bags/app/repositories"
)

type Auth interface {
	SignIn(request request.CreateSignInRequest) (accessToken string, err error)
}

type AuthService struct {
	AuthRepository repositories.Auth
}

// SignIn implements Auth.
func (svc *AuthService) SignIn(request request.CreateSignInRequest) (accessToken string, err error) {
	return "", nil
}

func NewAuthService(authRepository repositories.Auth) Auth {
	return &AuthService{
		AuthRepository: authRepository,
	}
}
