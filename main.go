package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/KaiGartenschlaeger/go-webserver/handler"
)

const port = 8080

func main() {
	http.HandleFunc("/home", handler.HomeHandler)
	http.HandleFunc("/about", handler.AboutHandler)
	http.HandleFunc("/health", handler.HealthHandler)
	http.HandleFunc("/", handler.CatchAllHandler)

	log.Printf("Starting server on PORT %d", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	must(err)
}
