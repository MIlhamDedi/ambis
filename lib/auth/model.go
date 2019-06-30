package auth

import "github.com/dgrijalva/jwt-go"

type (
	Authorization struct {
		AppName string `json:"appname"`
		Role    string `json:"role"`
	}

	Authentication struct {
		ID             string          `json:"username"`
		Authorizations []Authorization `json:"authorizations"`
	}

	Claims struct {
		Authentication Authentication
		jwt.StandardClaims
	}
)
