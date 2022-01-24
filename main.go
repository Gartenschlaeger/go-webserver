package main

import (
	"fmt"
	"log"
	"net/http"

	"webserver/handler"
	"webserver/helper"
)

const port = 8080

func main() {
	http.HandleFunc("/", handler.StaticHandler)
	http.HandleFunc("/home", handler.HomeHandler)
	http.HandleFunc("/about", handler.AboutHandler)
	http.HandleFunc("/ping", handler.PingHandler)

	log.Printf("Starting server on PORT %d", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	helper.Must(err)
}
