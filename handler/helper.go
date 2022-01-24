package handler

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func logRequest(r *http.Request) {
	log.Printf("%v %v %v\n", r.Proto, r.Method, r.URL.Path)
}

func renderTemplate(templateName string, model interface{}, w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles(fmt.Sprintf("templates/%s.html", templateName))
	if err != nil {
		log.Println("Failed to parse template", err)
	} else {
		template.Execute(w, model)
	}
}
