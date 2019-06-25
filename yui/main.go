package main

import (
	"ambis/lib/config"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

const (
	defaultPortAddr = 5001
)

type loginPageHandler struct{}

func (loginPageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./yui/template/login.html")
	if err != nil {
		fmt.Println("template not found")
		return
	}
	tmpl.Execute(w, nil)
}

func main() {
	fmt.Println("This is Yui")
	portAddr := config.GetPortAddr(defaultPortAddr)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./yui/static"))))
	http.Handle("/login/", loginPageHandler{})

	fmt.Printf("Yui is serving on %s...\n", portAddr)
	log.Fatal(http.ListenAndServe(portAddr, nil))
}
