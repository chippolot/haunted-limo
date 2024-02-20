package api

import (
	"html/template"
	"net/http"

	api "github.com/chippolot/haunted-limo/api/_pkg"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmplPath := api.GetTemplatePath("index.html")
	tmpl := template.Must(template.ParseFiles(tmplPath))
	tmpl.Execute(w, nil)
}
