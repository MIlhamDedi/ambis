package main

import (
	"ambis/lib/auth"
	"ambis/lib/config"
	"ambis/yui/handler"
	"ambis/yui/pb"
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("This is Yui")
	appConfig := config.Get(config.YUI)
	s := NewServer(*appConfig)
	log.Fatal(s.ListenAndServe())
}

func NewServer(appConfig config.Config) *http.Server {
	portAddr := appConfig.PortAddr
	authService := auth.AuthServiceImpl{SigningSecret: appConfig.SigningSecret}
	conn, err := grpc.Dial(appConfig.KiritoEndpoint, grpc.WithInsecure())
	if err != nil {
		log.Panic(err)
	}

	userServiceClient := pb.NewUserServiceClient(conn)

	fmt.Printf("Yui is serving on %s...\n", portAddr)
	return &http.Server{
		Addr:           portAddr,
		Handler:        handler.New(userServiceClient, &authService),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
