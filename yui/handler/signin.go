package handler

import (
	"ambis/lib/auth"
	"ambis/yui/account"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

/// Signin Handler
// The Following Contains Signin Handler Implementation
type Signin struct {
	UserService account.UserService
	AuthService auth.AuthService
}

func (h Signin) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.handlePost(w, r)
	default:
		h.handleGet(w, r)
	}
}

func (h Signin) handleGet(w http.ResponseWriter, r *http.Request) error {
	tmpl, err := template.ParseFiles("./yui/template/signin.html")
	if err != nil {
		fmt.Println("template not found")
		return err
	}
	tmpl.Execute(w, nil)
	return nil
}

func (h Signin) handlePost(w http.ResponseWriter, r *http.Request) error {
	verifyUser := &account.VerifyUser{}
	err := json.NewDecoder(r.Body).Decode(verifyUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	_, err = h.UserService.VerifyUser(verifyUser)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return err
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	authentication := auth.Authentication{ID: verifyUser.Username}
	createToken := auth.CreateToken{
		Authentication: authentication,
		ExpirationTime: expirationTime.Unix(),
	}

	token, err := h.AuthService.CreateToken(&createToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "ambis-session",
		Value:   token,
		Expires: expirationTime,
	})

	return nil
}
