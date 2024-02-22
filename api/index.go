package api

import (
	"html/template"
	"net/http"

	"github.com/chippolot/haunted-limo/api/_pkg/common"
	"github.com/chippolot/haunted-limo/api/_pkg/data"
	viewmodels "github.com/chippolot/haunted-limo/api/_pkg/view_models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// Load story data
	storyDataList, err := data.LoadStoryData()
	if err != nil {
		http.Error(w, "Failed to load story data", http.StatusBadRequest)
		return
	}

	model := viewmodels.IndexModel{
		Stories: storyDataList,
	}

	tmplPath := common.GetDataFilePath("templates/index.gohtml")
	tmpl := template.Must(template.ParseFiles(tmplPath))
	if err := tmpl.Execute(w, model); err != nil {
		panic(err)
	}
}
