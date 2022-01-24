package handler

import (
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

	helper.HandleTemplateRequest("home", model, w, r)
}
