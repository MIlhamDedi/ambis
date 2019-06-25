package main

import (
	"ambis/lib/config"
	"fmt"
	"log"
	"net/http"
)

const (
	defaultPortAddr = 3001
)

func main() {
	fmt.Println("This is Asuna")
	portAddr := config.GetPortAddr(defaultPortAddr)

	http.Handle("/", http.FileServer(http.Dir("./asuna/build")))

	fmt.Printf("Asuna is serving on %s...\n", portAddr)
	log.Fatal(http.ListenAndServe(portAddr, nil))
}
