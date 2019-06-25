package main

import (
	"fmt"
	"log"
	"net/http"

	"ambis/aincrad/proxy"
	"ambis/lib/config"
)

const (
	defaultPortAddr = 5000
)

func main() {
	fmt.Println("This is Aincrad")
	portAddr := config.GetPortAddr(defaultPortAddr)

	proxyHandler := proxy.NewHandler()
	http.Handle("/", proxyHandler)

	fmt.Printf("Aincrad is serving on %s...\n", portAddr)
	log.Fatal(http.ListenAndServe(portAddr, nil))
}
