package api

import (
	"html/template"
	"net/http"

	api "github.com/chippolot/haunted-limo/api/_pkg"
	blunders "github.com/chippolot/haunted-limo/api/blunders/_pkg"
)

func Blunders(w http.ResponseWriter, r *http.Request) {
	// Prep data provider
	connectionString := blunders.GetMySQLConnectionString()
	dataProvider := blunders.MakeSQLDataProvider(connectionString)
	defer dataProvider.Close()

	// Get most recent story
	result, err := dataProvider.GetMostRecentStory()
	if err != nil {
		panic(err)
	}

	tmplPath := api.GetTemplatePath("blunders.html")
	tmpl := template.Must(template.ParseFiles(tmplPath))
	tmpl.Execute(w, result)
}
