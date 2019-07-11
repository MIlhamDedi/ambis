package auth

import (
	"ambis/lib/base"
	"ambis/lib/constants"

	"github.com/dgrijalva/jwt-go"
)

type (
	AuthService interface {
		CreateToken(cmd *CreateToken) (string, error)
		VerifyToken(cmd *VerifyToken) (Authentication, error)
	}

	AuthServiceImpl struct {
		SigningSecret string
	}
)

func (s *AuthServiceImpl) CreateToken(cmd *CreateToken) (string, error) {
	claims := Claims{
		Authentication: cmd.Authentication,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: cmd.ExpirationTime,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.SigningSecret))
	if err != nil {
		return "", constants.ErrInternalError
	}

	return tokenString, nil
}

func (s *AuthServiceImpl) VerifyToken(cmd *VerifyToken) (Authentication, error) {
	authentication := Authentication{}
	claims := Claims{}
	token, err := jwt.ParseWithClaims(cmd.Token, claims, func(token *jwt.Token) (interface{}, error) {
		return s.SigningSecret, nil
	})
	if !token.Valid {
		return authentication, constants.ErrExpiredToken
	}
	if err != nil {
		return authentication, err
	}

	authentication = claims.Authentication
	return authentication, nil
}

func NewService(base *base.Base) (*AuthServiceImpl, error) {
	return &AuthServiceImpl{
		SigningSecret: base.Config.SigningSecret,
	}, nil
}
