package handler

import (
	"log"
	"net/http"
	"path"
)

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("StaticHandler %v %v %v\n", r.Proto, r.Method, r.URL.Path)

	if r.URL.Path == "/" {
		HomeHandler(w, r)
	} else {
		path := path.Join("wwwroot", r.URL.Path[1:])
		http.ServeFile(w, r, path)
	}
}
