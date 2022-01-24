package handler

import (
	"log"
	"net/http"

	"webserver/helper"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("AboutHandler %v %v %v\n", r.Proto, r.Method, r.URL.Path)

	helper.HandleTemplateRequest("about", nil, w, r)
}
