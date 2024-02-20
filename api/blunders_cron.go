package api

import (
	"net/http"

	common "github.com/chippolot/haunted-limo/api/_pkg"
	blunders "github.com/chippolot/haunted-limo/api/_pkg/blunders"
	"github.com/chippolot/jokegen"
)

func Cron(w http.ResponseWriter, r *http.Request) {
	// Prep data provider
	connectionString := common.GetMySQLConnectionString()
	dataProvider := blunders.MakeSQLDataProvider(connectionString)
	defer dataProvider.Close()

	// Get most recent story
	openAIToken := common.GetOpenAIToken()
	options := jokegen.StoryOptions{ForceRegenerate: true}
	_, err := jokegen.GenerateStory(openAIToken, jokegen.Misunderstanding, dataProvider, options)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
}
