package handler

import (
	"ambis/lib/auth"
	"ambis/yui/account"
	"net/http"
)

// New Yui Root Handler requires auth.AuthService and account.UserService
func New(authService auth.AuthService, userService account.UserService) http.Handler {
	Signin := Signin{UserService: userService, AuthService: authService}
	Signup := Signup{UserService: userService}

	mux := http.NewServeMux()
	mux.Handle("/signin/", http.StripPrefix("/signin/", Signin))
	mux.Handle("/signup/", http.StripPrefix("/signup/", Signup))

	return mux
}
