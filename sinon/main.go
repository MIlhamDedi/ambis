package main

import (
	"ambis/lib/config"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("This is Sinon")
	appConfig := config.Get(config.SINON)
	portAddr := appConfig.PortAddr

	http.Handle("/", http.FileServer(http.Dir("./sinon/build")))

	fmt.Printf("Sinon is serving on %s...\n", portAddr)
	log.Fatal(http.ListenAndServe(portAddr, nil))
}
