package main

import (
	"fmt"

	"ambis/kirito/account"
	"ambis/kirito/pb"
	"ambis/lib/config"
	"ambis/lib/utils/db"
	"net"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	fmt.Println("This is Kirito")
	appConfig := config.Get(config.KIRITO)
	db := db.New(appConfig)

	userService, err := account.NewUserService(db)
	if err != nil {
		log.Panic(err)
	}

	listener, err := net.Listen("tcp", appConfig.PortAddr)
	if err != nil {
		log.Panic(err)
	}

	srv := grpc.NewServer()
	pb.RegisterUserServiceServer(srv, userService)
	reflection.Register(srv)

	fmt.Printf("Kirito is serving on %s...\n", appConfig.PortAddr)
	if err = srv.Serve(listener); err != nil {
		log.Panic("heree")
		log.Panic(err)
	}
}
