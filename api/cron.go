package api

import (
	"fmt"
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

	// Check for required story type param
	storyTypeParams, ok := r.URL.Query()["storyType"]
	if !ok || len(storyTypeParams) != 1 || len(storyTypeParams[0]) < 1 {
		http.Error(w, "Expected single required query parameter: storyType", http.StatusBadRequest)
		return
	}
	storyTypeParam := storyTypeParams[0]

	// Validate story type
	storyType, err := jokegen.ParseStoryType(storyTypeParam)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unknown story type: %s", storyTypeParam), http.StatusBadRequest)
		return
	}

	// Prep data provider
	connectionString := common.GetMySQLConnectionString()
	dataProvider := common.MakeSQLDataProvider(connectionString)
	defer dataProvider.Close()

	// Generate stories
	openAIToken := common.GetOpenAIToken()
	options := jokegen.StoryOptions{ForceRegenerate: true}
	_, err = jokegen.GenerateStory(openAIToken, storyType, dataProvider, options)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
}
