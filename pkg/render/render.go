package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var functions = template.FuncMap{}

func RenderTemplate(w http.ResponseWriter, r *http.Request, templateName string, model interface{}) {
	templateCache, errTemplateCache := createTemplateCache()
	if errTemplateCache != nil {
		log.Fatalln(errTemplateCache)
	}

	template, ok := templateCache[fmt.Sprintf("%v.html", templateName)]
	if !ok {
		log.Fatal("template not ok")
	}

	buf := new(bytes.Buffer)
	_ = template.Execute(buf, model)

	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println("Failed to write template to buffer", errTemplateCache)
	}
}

func createTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		log.Println("Page is currently", name)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return nil, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return nil, err
			}

			cache[name] = ts
		}
	}

	return cache, nil
}
