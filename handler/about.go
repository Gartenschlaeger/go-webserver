package handler

import (
	"net/http"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	renderTemplate("about", nil, w, r)
}
