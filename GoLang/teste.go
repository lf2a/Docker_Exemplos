package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type home struct {
	Titulo  string
	Request string
}

func helloWorld(res http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))

	h := home{
		Titulo: "Testando Golang e Docker",
	}

	fmt.Println(req)

	tmpl.Execute(res, h)
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":3000", nil)
}
