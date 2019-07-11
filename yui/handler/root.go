package handler

import (
	"ambis/lib/base"
	"ambis/yui/handler/signin"
	"ambis/yui/handler/signup"
	"net/http"
)

// New Yui Root Handler requires auth.AuthService and account.UserService
func New(b *base.Base) http.Handler {
	Signin, err := signin.New(b)
	if err != nil {
		b.Log.Panic(Signin)
	}
	Signup, err := signup.New(b)
	if err != nil {
		b.Log.Panic(Signin)
	}

	mux := http.NewServeMux()
	mux.Handle("/sso/", http.StripPrefix("/sso/", Signin))
	mux.Handle("/signup/", http.StripPrefix("/signup/", Signup))

	return mux
}
