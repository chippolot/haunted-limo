package api

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmplPath, _ := filepath.Abs("api/data/templates/index.html")
	tmpl := template.Must(template.ParseFiles(tmplPath))
	tmpl.Execute(w, nil)
}
