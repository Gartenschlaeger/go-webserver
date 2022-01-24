package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"

	"webserver/models"
)

func staticContentHandler(response http.ResponseWriter, request *http.Request) {
	path := path.Join("wwwroot", request.URL.Path[1:])
	fmt.Printf("staticContentHandler path %v\n", path)

	http.ServeFile(response, request, path)
}

func templateHandler(response http.ResponseWriter, request *http.Request) {
	model := models.IndexModel{
		Title: "Test"}

	template, err := template.ParseFiles("wwwroot/index.html")
	must(err)

	template.Execute(response, model)
}

func pingHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "PONG")
}

func main() {
	http.HandleFunc("/", staticContentHandler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/template", templateHandler)

	http.ListenAndServe(":8080", nil)
}
