package main

import (
	"ambis/lib/base"
	"ambis/lib/config"
	"ambis/yui/handler"
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("This is Yui")
	appConfig := config.Get(config.YUI)
	base, err := base.New(appConfig)
	if err != nil {
		log.Panic(err)
	}
	s := NewServer(base)
	log.Fatal(s.ListenAndServe())
}

func NewServer(b *base.Base) *http.Server {
	portAddr := b.Config.PortAddr

	fmt.Printf("Yui is serving on %s...\n", portAddr)
	return &http.Server{
		Addr:           portAddr,
		Handler:        handler.New(b),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
