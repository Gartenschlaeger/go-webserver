package handlers

import (
	"net/http"

	"github.com/KaiGartenschlaeger/go-webserver/pkg/logging"
	"github.com/KaiGartenschlaeger/go-webserver/pkg/render"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	logging.LogRequest(r)

	render.RenderTemplate("about", nil, w, r)
}
