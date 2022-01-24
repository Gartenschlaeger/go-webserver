package handler

import (
	"html/template"
	"log"
	"net/http"

	"webserver/helper"
	"webserver/models"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("HomeHandler %v %v %v\n", r.Proto, r.Method, r.URL.Path)

	model := models.IndexModel{
		Title:   "Test",
		Message: "Welcome to the home page"}

	template, err := template.ParseFiles("templates/home.html")
	helper.Must(err)

	template.Execute(w, model)
}
