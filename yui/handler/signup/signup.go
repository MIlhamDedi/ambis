package signup

import (
	"ambis/lib/base"
	"ambis/yui/pb"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"google.golang.org/grpc"
)

// User API specification (temp until grpc-web implementation)
type User struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// Signup Handler
// The Following Contains SignUp Handler Implementation
type Signup struct {
	Base              *base.Base
	UserServiceClient pb.UserServiceClient
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
	user := &User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	pbUser := pb.User{
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
	}

	_, err = h.UserServiceClient.Create(context.Background(), &pbUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	return nil
}

func New(b *base.Base) (*Signup, error) {
	conn, err := grpc.Dial(b.Config.KiritoEndpoint, grpc.WithInsecure())
	if err != nil {
		b.Log.Panic(err)
	}
	userServiceClient := pb.NewUserServiceClient(conn)
	return &Signup{
		Base:              b,
		UserServiceClient: userServiceClient,
	}, nil
}
