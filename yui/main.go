package main

import (
	"ambis/lib/auth"
	"ambis/lib/config"
	"ambis/yui/account"
	"ambis/yui/handler"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("This is Yui")
	appConfig := config.Get(config.YUI)
	s := NewServer(*appConfig)
	log.Fatal(s.ListenAndServe())
}

func NewServer(appConfig config.Config) *http.Server {
	portAddr := appConfig.PortAddr
	db, err := sql.Open(appConfig.DBMode, fmt.Sprintf("%s:%s@/%s", appConfig.DBUser, appConfig.DBPassword, appConfig.DBName))
	if err != nil {
		log.Fatal(err)
	}

	userRepo := account.UserRepoDB{DB: db}
	userService := account.UserServiceImpl{UserRepo: &userRepo}
	authService := auth.AuthServiceImpl{SigningSecret: appConfig.SigningSecret}

	fmt.Printf("Yui is serving on %s...\n", portAddr)
	return &http.Server{
		Addr:           portAddr,
		Handler:        handler.New(&authService, &userService),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
