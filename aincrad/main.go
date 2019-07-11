package main

import (
	"fmt"
	"log"
	"net/http"

	"ambis/aincrad/proxy"
	"ambis/lib/base"
	"ambis/lib/config"
)

func main() {
	fmt.Println("This is Aincrad")
	appConfig := config.Get(config.AINCRAD)
	base, err := base.New(appConfig)
	if err != nil {
		log.Panic(err)
	}
	portAddr := appConfig.PortAddr

	proxyHandler := proxy.New(base)
	http.Handle("/", proxyHandler)

	fmt.Printf("Aincrad is serving on %s...\n", portAddr)
	log.Fatal(http.ListenAndServe(portAddr, nil))
}
