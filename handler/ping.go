package handler

import (
	"fmt"
	"log"
	"net/http"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("PingHandler %v %v %v\n", r.Proto, r.Method, r.URL.Path)

	fmt.Fprint(w, "PONG")
}
