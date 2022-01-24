package handler

import (
	"net/http"
	"webserver/helper"
	"webserver/models"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	helper.LogRequest(r)

	model := models.IndexModel{
		Title:   "Test",
		Message: "Welcome to the home page"}

	helper.HandleTemplateRequest("home", model, w, r)
}
