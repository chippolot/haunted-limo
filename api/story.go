package api

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/chippolot/haunted-limo/api/_pkg/common"
	"github.com/chippolot/haunted-limo/api/_pkg/data"
	viewmodels "github.com/chippolot/haunted-limo/api/_pkg/view_models"
	"github.com/chippolot/jokegen"
)

func Story(w http.ResponseWriter, r *http.Request) {
	// Load story data
	storyDataList, err := data.LoadStoryData()
	if err != nil {
		http.Error(w, "Failed to load story data", http.StatusBadRequest)
		return
	}

	// Check for required story type param
	storyTypeParams, ok := r.URL.Query()["storyType"]
	if !ok || len(storyTypeParams) != 1 || len(storyTypeParams[0]) < 1 {
		http.Error(w, "Expected single required query parameter: storyType", http.StatusBadRequest)
		return
	}
	storyTypeParam := storyTypeParams[0]

	// Find story data
	storyData, err := data.FindStoryData(storyDataList, storyTypeParam)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to load story data: %s", storyTypeParam), http.StatusBadRequest)
		panic(err)
	}

	// Validate story type
	storyType, err := jokegen.ParseStoryType(storyTypeParam)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unknown story type: %s", storyTypeParam), http.StatusBadRequest)
		return
	}

	// Prep data provider
	connectionString := common.GetMySQLConnectionString()
	dataProvider := data.MakeSQLDataProvider(connectionString)
	defer dataProvider.Close()

	// Get most recent story
	result, err := dataProvider.GetMostRecentStory(storyType)
	if err != nil {
		panic(err)
	}

	model := viewmodels.StoryModel{
		Cfg:   storyData,
		Story: result.Story,
	}

	tmplPath := common.GetDataFilePath("templates/story.gohtml")
	tmpl := template.Must(template.ParseFiles(tmplPath))
	if err := tmpl.Execute(w, model); err != nil {
		panic(err)
	}
}
