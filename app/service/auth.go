package service

import (
	"github.com/bagusyanuar/go-olin-bags/app/config"
	"github.com/bagusyanuar/go-olin-bags/app/http/request"
	"github.com/bagusyanuar/go-olin-bags/app/repositories"
	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/bagusyanuar/go-olin-bags/model"
)

type Auth interface {
	SignIn(request request.CreateSignInRequest) (accessToken string, err error)
}

type AuthService struct {
	JWTConfig      config.JWT
	AuthRepository repositories.Auth
}

// SignIn implements Auth.
func (svc *AuthService) SignIn(request request.CreateSignInRequest) (accessToken string, err error) {
	user := model.User{
		Username: request.Username,
		Password: &request.Password,
	}
	u, e := svc.AuthRepository.SignIn(user)
	if err != nil {
		return "", e
	}
	jwtSign := common.JWTSignReturn{
		UserID:   u.ID,
		Username: u.Username,
	}
	return common.CreateAccessToken(&svc.JWTConfig, &jwtSign)
}

func NewAuthService(authRepository repositories.Auth, jwtConfig config.JWT) Auth {
	return &AuthService{
		JWTConfig:      jwtConfig,
		AuthRepository: authRepository,
	}
}
