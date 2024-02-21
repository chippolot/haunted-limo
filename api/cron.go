package api

import (
	"net/http"
	"os"
	"strings"

	common "github.com/chippolot/haunted-limo/api/_pkg"
	"github.com/chippolot/jokegen"
)

func Cron(w http.ResponseWriter, r *http.Request) {
	// Get the bearer token from the Authorization header
	authHeader := r.Header.Get("Authorization")
	splitToken := strings.Split(authHeader, "Bearer ")
	var authToken string
	if len(splitToken) > 1 {
		authToken = splitToken[1]
	}

	// Check if the token is not found or does not match the CRON_SECRET
	if authToken == "" || authToken != os.Getenv("CRON_SECRET") {
		http.Error(w, `{"error": "Unauthorized"}`, http.StatusUnauthorized)
		return
	}

	// Prep data provider
	connectionString := common.GetMySQLConnectionString()
	dataProvider := common.MakeSQLDataProvider(connectionString)
	defer dataProvider.Close()

	openAIToken := common.GetOpenAIToken()

	// Generate stories
	options := jokegen.StoryOptions{ForceRegenerate: true}
	storyTypes := []jokegen.StoryType{jokegen.Misunderstanding, jokegen.Slapstick, jokegen.Hex}
	for _, storyType := range storyTypes {
		_, err := jokegen.GenerateStory(openAIToken, storyType, dataProvider, options)
		if err != nil {
			panic(err)
		}
	}

	w.WriteHeader(http.StatusOK)
}
