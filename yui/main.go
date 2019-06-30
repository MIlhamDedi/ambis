package main

import (
	"ambis/lib/auth"
	"ambis/lib/config"
	"ambis/yui/account"
	"ambis/yui/handler"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const (
	defaultPortAddr = 3000
)

func main() {
	fmt.Println("This is Yui")
	portAddr := config.GetPortAddr(defaultPortAddr)

	db, err := sql.Open(config.DBMode, fmt.Sprintf("%s:%s@/%s", config.DBUser, config.DBPassword, config.DBName))
	if err != nil {
		log.Fatal(err)
	}

	userRepo := account.UserRepoDB{DB: db}
	userService := account.UserServiceImpl{UserRepo: &userRepo}

	authService := auth.AuthServiceImpl{SigningSecret: config.SigningSecret}

	signinHandler := handler.Signin{UserService: &userService, AuthService: &authService}
	signupHandler := handler.Signup{UserService: &userService}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./yui/static"))))
	http.Handle("/signin/", http.StripPrefix("/signin/", signinHandler))
	http.Handle("/signup/", http.StripPrefix("/signup/", signupHandler))

	fmt.Printf("Yui is serving on %s...\n", portAddr)
	log.Fatal(http.ListenAndServe(portAddr, nil))
}
