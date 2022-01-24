package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func RenderTemplate(templateName string, model interface{}, w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles(fmt.Sprintf("templates/%s.html", templateName))
	if err != nil {
		log.Println("Failed to parse template", err)
	} else {
		template.Execute(w, model)
	}
}
