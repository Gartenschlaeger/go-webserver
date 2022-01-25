package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/KaiGartenschlaeger/go-webserver/pkg/config"
	"github.com/KaiGartenschlaeger/go-webserver/pkg/handlers"
	"github.com/KaiGartenschlaeger/go-webserver/pkg/render"
)

const port = 8080

func must(err error) {
	if err != nil {
		log.Panicln(err)
		panic(err)
	}
}

func main() {
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalln("Failed to create template cache", err)
	}

	app := config.AppConfig{}
	app.TemplateCache = tc

	render.NewTemplates(&app)

	http.HandleFunc("/home", handlers.HomeHandler)
	http.HandleFunc("/about", handlers.AboutHandler)
	http.HandleFunc("/health", handlers.HealthHandler)
	http.HandleFunc("/", handlers.CatchAllHandler)

	log.Printf("Starting server on PORT %d", port)

	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
