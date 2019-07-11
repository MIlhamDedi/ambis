package main

import (
	"fmt"

	"ambis/kirito/account"
	"ambis/kirito/pb"
	"ambis/lib/base"
	"ambis/lib/config"
	"net"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	fmt.Println("This is Kirito")
	appConfig := config.Get(config.KIRITO)
	base, err := base.New(appConfig)
	if err != nil {
		log.Panic(err)
	}

	listener, err := net.Listen("tcp", appConfig.PortAddr)
	if err != nil {
		log.Panic(err)
	}

	srv, err := NewServer(base)
	if err != nil {
		base.Log.Panic(err)
	}

	if err = srv.Serve(listener); err != nil {
		log.Panic(err)
	}
}

func NewServer(b *base.Base) (*grpc.Server, error) {
	userService, err := account.NewService(b)
	if err != nil {
		log.Panic(err)
	}

	srv := grpc.NewServer()
	pb.RegisterUserServiceServer(srv, userService)
	reflection.Register(srv)

	fmt.Printf("Kirito is serving on %s...\n", b.Config.PortAddr)
	return srv, nil
}
