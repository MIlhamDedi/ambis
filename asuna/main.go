package main

import (
	"ambis/lib/config"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("This is Asuna")
	appConfig := config.Get(config.ASUNA)
	portAddr := appConfig.PortAddr

	http.Handle("/", http.FileServer(http.Dir("./asuna/build")))

	fmt.Printf("Asuna is serving on %s...\n", portAddr)
	log.Fatal(http.ListenAndServe(portAddr, nil))
}
