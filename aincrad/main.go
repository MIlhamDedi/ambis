package main

import (
	"fmt"
	"log"
	"net/http"

	"ambis/aincrad/proxy"
)

func main() {
	fmt.Println("This is Aincrad")

	proxyHandler := proxy.NewHandler()
	http.Handle("/", proxyHandler)

	fmt.Println("Aincrad is serving on port 5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
