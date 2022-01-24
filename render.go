package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func renderTemplate(templateName string, model interface{}, w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles(fmt.Sprintf("templates/%s.html", templateName))
	if err != nil {
		log.Println("Failed to parse template", err)
	} else {
		template.Execute(w, model)
	}
}
