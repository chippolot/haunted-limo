package api

import (
	"net/http"

	"github.com/chippolot/blunder"
	blunders "github.com/chippolot/haunted-limo/api/blunders/_pkg"
)

func Cron(w http.ResponseWriter, r *http.Request) {
	// Prep data provider
	connectionString := blunders.GetMySQLConnectionString()
	dataProvider := blunders.MakeSQLDataProvider(connectionString)
	defer dataProvider.Close()

	// Get most recent story
	openAIToken := blunders.GetOpenAIToken()
	_, err := blunder.GenerateStory(openAIToken, dataProvider, blunder.StoryOptions{ForceRegenerate: true})
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
}
