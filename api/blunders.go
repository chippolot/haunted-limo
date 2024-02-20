package handler

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/chippolot/blunder"
	"github.com/chippolot/haunted-limo/api/internal/blunders"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// Resolve API key
	token := os.Getenv("OPEN_AI_API_KEY")
	if token == "" {
		panic("OpenAI API key not found in environment variables")
	}

	dsn := os.Getenv("DSN")
	if dsn == "" {
		panic("DSN not found in environment variables")
	}

	dataProvider := blunders.MakeSQLDataProvider(dsn)

	// Generate Story
	options := blunder.StoryOptions{}
	result, err := blunder.GenerateStory(token, dataProvider, options)
	if err != nil {
		panic(err)
	}

	// Write JSON response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		panic(err)
	}
}
