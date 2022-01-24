package handler

import (
	"net/http"
	"webserver/models"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r)

	model := models.IndexModel{
		Title:   "Test",
		Message: "Welcome to the home page"}

	renderTemplate("home", model, w, r)
}
