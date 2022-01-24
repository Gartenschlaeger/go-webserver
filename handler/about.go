package handler

import (
	"net/http"
	"webserver/helper"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	helper.LogRequest(r)
	helper.HandleTemplateRequest("about", nil, w, r)
}
