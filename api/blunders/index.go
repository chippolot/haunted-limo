package api

import (
	"html/template"
	"net/http"

	"github.com/chippolot/haunted-limo/api/_pkg/blunders"
)

func Blunders(w http.ResponseWriter, r *http.Request) {
	// Prep data provider
	connectionString := getMySQLConnectionString()
	dataProvider := blunders.MakeSQLDataProvider(connectionString)
	defer dataProvider.Close()

	// Get most recent story
	result, err := dataProvider.GetMostRecentStory()
	if err != nil {
		panic(err)
	}

	tmpl := template.Must(template.ParseFiles("templates/blunders.html"))
	tmpl.Execute(w, result)
}
