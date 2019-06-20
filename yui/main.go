package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"ambis/lib/config"
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
	rootPath := config.YuiRootPath

	http.Handle(rootPath+"static/", http.StripPrefix(rootPath+"static/", http.FileServer(http.Dir("./yui/static"))))
	http.Handle(rootPath, loginPageHandler{})

	fmt.Println("Yui is serving on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
