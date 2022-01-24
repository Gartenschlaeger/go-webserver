package logging

import (
	"log"
	"net/http"
)

func LogRequest(r *http.Request) {
	log.Printf("%v %v %v\n", r.Proto, r.Method, r.URL.Path)
}
