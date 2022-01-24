package handlers

import (
	"net/http"

	"github.com/KaiGartenschlaeger/go-webserver/pkg/logging"
	"github.com/KaiGartenschlaeger/go-webserver/pkg/models"
	"github.com/KaiGartenschlaeger/go-webserver/pkg/render"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	logging.LogRequest(r)

	model := models.IndexModel{
		Title:   "Test",
		Message: "Welcome to the home page"}

	render.RenderTemplate("home", model, w, r)
}
