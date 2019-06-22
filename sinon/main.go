package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("This is Sinon")

	http.Handle("/", http.FileServer(http.Dir("./sinon/build")))

	fmt.Println("Sinon is serving on port 3002...")
	log.Fatal(http.ListenAndServe(":3002", nil))
}
