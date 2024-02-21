package api

import (
	"net/http"

	common "github.com/chippolot/haunted-limo/api/_pkg"
	"github.com/chippolot/jokegen"
)

func Cron(w http.ResponseWriter, r *http.Request) {
	// Prep data provider
	connectionString := common.GetMySQLConnectionString()
	dataProvider := common.MakeSQLDataProvider(connectionString)
	defer dataProvider.Close()

	openAIToken := common.GetOpenAIToken()

	// Generate stories
	options := jokegen.StoryOptions{ForceRegenerate: true}
	storyTypes := []jokegen.StoryType{jokegen.Misunderstanding, jokegen.Slapstick}
	for _, storyType := range storyTypes {
		_, err := jokegen.GenerateStory(openAIToken, storyType, dataProvider, options)
		if err != nil {
			panic(err)
		}
	}

	w.WriteHeader(http.StatusOK)
}
