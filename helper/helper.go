package helper

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Must(err error) {
	if err != nil {
		log.Panicln(err)
		panic(err)
	}
}

func HandleTemplateRequest(templateName string, model interface{}, w http.ResponseWriter, r *http.Request) {
	log.Printf("handleTemplateRequest %v %v %v\n", r.Proto, r.Method, r.URL.Path)

	template, err := template.ParseFiles(fmt.Sprintf("templates/%s.html", templateName))
	Must(err)

	template.Execute(w, model)
}
