package handler

import (
	"html/template"
	"log"
	"net/http"
	"webserver/helper"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("AboutHandler %v %v %v\n", r.Proto, r.Method, r.URL.Path)

	template, err := template.ParseFiles("templates/about.html")
	helper.Must(err)

	template.Execute(w, nil)
}
