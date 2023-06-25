package productionhouse

import (
	request "github.com/bagusyanuar/go-olin-bags/app/http/request/production-house"
	repository "github.com/bagusyanuar/go-olin-bags/app/repositories/production-house"
	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/bagusyanuar/go-olin-bags/config"
	"github.com/bagusyanuar/go-olin-bags/model"
	"golang.org/x/crypto/bcrypt"
)

type Auth interface {
	SignIn(request request.SignInRequest) (string, error)
}

type AuthService struct {
	JWTConfig      config.JWT
	AuthRepository repository.Auth
}

// SignIn implements Auth.
func (svc *AuthService) SignIn(request request.SignInRequest) (string, error) {
	user := model.User{
		Username: request.Username,
		Password: &request.Password,
	}

	u, e := svc.AuthRepository.SignIn(user)
	if e != nil {
		return "", e
	}

	err := bcrypt.CompareHashAndPassword([]byte(*u.Password), []byte(request.Password))

	if err != nil {
		return "", common.ErrorPasswordNotMatch
	}
	jwtSign := common.JWTSignReturn{
		UserID:   u.ID,
		Username: u.Username,
	}

	return common.CreateAccessToken(&svc.JWTConfig, &jwtSign)
}

func NewAuthService(authRepository repository.Auth, config config.JWT) Auth {
	return &AuthService{
		AuthRepository: authRepository,
		JWTConfig:      config,
	}
}
