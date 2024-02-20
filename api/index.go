package api

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("api/data/templates/index.html"))
	tmpl.Execute(w, nil)
}
