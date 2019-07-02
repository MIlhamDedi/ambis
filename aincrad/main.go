package main

import (
	"fmt"
	"log"
	"net/http"

	"ambis/aincrad/proxy"
	"ambis/lib/config"
)

func main() {
	fmt.Println("This is Aincrad")
	appConfig := config.Get(config.AINCRAD)
	portAddr := appConfig.PortAddr

	proxyHandler := proxy.NewHandler()
	http.Handle("/", proxyHandler)

	fmt.Printf("Aincrad is serving on %s...\n", portAddr)
	log.Fatal(http.ListenAndServe(portAddr, nil))
}
