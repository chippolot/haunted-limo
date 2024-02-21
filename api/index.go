package api

import (
	"html/template"
	"net/http"
	"os"
)

func Index(w http.ResponseWriter, r *http.Request) {
	baseTmplDir := os.Getenv("BASE_TEMPLATE_DIR")
	tmplPath := baseTmplDir + "data/templates/index.gohtml"
	tmpl := template.Must(template.ParseFiles(tmplPath))
	tmpl.Execute(w, nil)
}
