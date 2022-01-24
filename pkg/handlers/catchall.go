package handlers

import (
	"net/http"
	"path"

	"github.com/KaiGartenschlaeger/go-webserver/pkg/logging"
)

func CatchAllHandler(w http.ResponseWriter, r *http.Request) {
	logging.LogRequest(r)

	if r.URL.Path == "/" {
		HomeHandler(w, r)
	} else {
		path := path.Join("wwwroot", r.URL.Path[1:])
		http.ServeFile(w, r, path)
	}
}
