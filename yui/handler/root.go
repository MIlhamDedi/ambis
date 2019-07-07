package handler

import (
	"ambis/lib/auth"
	"ambis/yui/pb"
	"net/http"
)

// User API specification (temp until grpc-web implementation)
type User struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// New Yui Root Handler requires auth.AuthService and account.UserService
func New(userServiceClient pb.UserServiceClient, authService auth.AuthService) http.Handler {
	Signin := Signin{UserServiceClient: userServiceClient, AuthService: authService}
	Signup := Signup{UserServiceClient: userServiceClient}

	mux := http.NewServeMux()
	mux.Handle("/sso/", http.StripPrefix("/sso/", Signin))
	mux.Handle("/signup/", http.StripPrefix("/signup/", Signup))

	return mux
}
