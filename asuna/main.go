package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./asuna/build")))

	fmt.Println("Asuna is serving on port 3001...")
	log.Fatal(http.ListenAndServe(":3001", nil))
}
