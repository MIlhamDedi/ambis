package handler

import (
	"ambis/lib/auth"
	"ambis/yui/pb"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

/// Signin Handler
// The Following Contains Signin Handler Implementation
type Signin struct {
	UserServiceClient pb.UserServiceClient
	AuthService       auth.AuthService
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
	tmpl, err := template.ParseFiles("./yui/template/sso.html")
	if err != nil {
		fmt.Println("template not found")
		return err
	}
	tmpl.Execute(w, nil)
	return nil
}

func (h Signin) handlePost(w http.ResponseWriter, r *http.Request) error {
	user := &User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	pbUser := pb.User{
		Username: user.Username,
		Password: user.Password,
	}

	_, err = h.UserServiceClient.Verify(context.Background(), &pbUser)
	if err != nil {
		log.Info(err)
		w.WriteHeader(http.StatusUnauthorized)
		return err
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	authentication := auth.Authentication{ID: user.Username}
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
		Path:    "/",
		Expires: expirationTime,
	})

	return nil
}
