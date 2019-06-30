package handler

import (
	"ambis/yui/account"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

// Signup Handler
// The Following Contains SignUp Handler Implementation
type Signup struct {
	UserService account.UserService
}

func (h Signup) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.handlePost(w, r)
	default:
		h.handleGet(w, r)
	}
}

func (h Signup) handleGet(w http.ResponseWriter, r *http.Request) error {
	tmpl, err := template.ParseFiles("./yui/template/signup.html")
	if err != nil {
		fmt.Println("template not found")
		return err
	}
	tmpl.Execute(w, nil)
	return nil
}

func (h Signup) handlePost(w http.ResponseWriter, r *http.Request) error {
	createUser := &account.CreateUser{}
	err := json.NewDecoder(r.Body).Decode(createUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	_, err = h.UserService.CreateUser(createUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	return nil
}
