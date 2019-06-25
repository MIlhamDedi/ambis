package main

import (
	"ambis/lib/config"
	"fmt"
	"log"
	"net/http"
)

const (
	defaultPortAddr = 3002
)

func main() {
	fmt.Println("This is Sinon")
	portAddr := config.GetPortAddr(defaultPortAddr)

	http.Handle("/", http.FileServer(http.Dir("./sinon/build")))

	fmt.Printf("Sinon is serving on %s...\n", portAddr)
	log.Fatal(http.ListenAndServe(portAddr, nil))
}
