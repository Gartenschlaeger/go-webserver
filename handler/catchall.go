package handler

import (
	"net/http"
	"path"
)

func CatchAllHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r)

	if r.URL.Path == "/" {
		HomeHandler(w, r)
	} else {
		path := path.Join("wwwroot", r.URL.Path[1:])
		http.ServeFile(w, r, path)
	}
}
