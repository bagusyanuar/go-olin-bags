package common

import (
	"github.com/bagusyanuar/go-olin-bags/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTClaims struct {
	jwt.RegisteredClaims
	UserID   uuid.UUID `json:"user_id"`
	Username string    `json:"username"`
}

type JWTSignReturn struct {
	UserID   uuid.UUID `json:"user_id"`
	Username string    `json:"username"`
}

func CreateAccessToken(cfg *config.JWT, jwtSign *JWTSignReturn) (accessToken string, err error) {
	JWTSigninMethod := jwt.SigningMethodHS256
	claims := JWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: cfg.Issuer,
		},
		UserID:   jwtSign.UserID,
		Username: jwtSign.Username,
	}
	token := jwt.NewWithClaims(JWTSigninMethod, claims)
	return token.SignedString([]byte(cfg.JWTSignatureKey))
}
