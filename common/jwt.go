package common

import (
	"github.com/bagusyanuar/go-olin-bags/config"
	"github.com/gin-gonic/gin"
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

func VerifyToken(token string, cfg *config.JWT) (*JWTClaims, error) {
	jwtSignaturKey := cfg.JWTSignatureKey
	t, e := jwt.ParseWithClaims(token, &JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtSignaturKey), nil
	})
	if e != nil {
		return nil, e
	}
	if claims, ok := t.Claims.(*JWTClaims); ok && t.Valid {
		return claims, nil
	}
	return nil, ErrorUnauthorized
}

func GetAuthorizedID(ctx *gin.Context) string {
	defer Catch(ctx)
	authorizedUser := ctx.MustGet("user").(*JWTClaims)
	return authorizedUser.UserID.String()
}
