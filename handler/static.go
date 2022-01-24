package handler

import (
	"net/http"
	"path"
	"webserver/helper"
)

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	helper.LogRequest(r)

	if r.URL.Path == "/" {
		HomeHandler(w, r)
	} else {
		path := path.Join("wwwroot", r.URL.Path[1:])
		http.ServeFile(w, r, path)
	}
}
