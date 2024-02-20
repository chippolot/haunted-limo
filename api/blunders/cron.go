package api

import (
	"net/http"

	"github.com/chippolot/blunder"
	"github.com/chippolot/haunted-limo/api/_pkg/blunders"
)

func Cron(w http.ResponseWriter, r *http.Request) {
	// Prep data provider
	connectionString := getMySQLConnectionString()
	dataProvider := blunders.MakeSQLDataProvider(connectionString)
	defer dataProvider.Close()

	// Get most recent story
	openAIToken := getOpenAIToken()
	_, err := blunder.GenerateStory(openAIToken, dataProvider, blunder.StoryOptions{ForceRegenerate: true})
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
}
