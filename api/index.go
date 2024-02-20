package api

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmplPath := "data/templates/index.html"
	tmpl := template.Must(template.ParseFiles(tmplPath))
	tmpl.Execute(w, nil)
}
