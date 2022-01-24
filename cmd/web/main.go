package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/KaiGartenschlaeger/go-webserver/pkg/handlers"
)

const port = 8080

func must(err error) {
	if err != nil {
		log.Panicln(err)
		panic(err)
	}
}

func main() {
	http.HandleFunc("/home", handlers.HomeHandler)
	http.HandleFunc("/about", handlers.AboutHandler)
	http.HandleFunc("/health", handlers.HealthHandler)
	http.HandleFunc("/", handlers.CatchAllHandler)

	log.Printf("Starting server on PORT %d", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	must(err)
}
