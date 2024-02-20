package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/chippolot/haunted-limo/api/_pkg/blunders"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Haunted Limo")
}

func blundersHandler(w http.ResponseWriter, r *http.Request) {
	// Resolve API key
	token := os.Getenv("OPEN_AI_API_KEY")
	if token == "" {
		panic("OpenAI API key not found in environment variables")
	}

	// Resolve DB connection string
	dsn := os.Getenv("DSN")
	if dsn == "" {
		panic("DSN not found in environment variables")
	}

	// Prep data provider
	dataProvider := blunders.MakeSQLDataProvider(dsn)
	defer dataProvider.Close()

	// Get most recent story
	result, err := dataProvider.GetMostRecentStory()
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

func router(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	switch path {
	case "/":
		indexHandler(w, r)
	case "/blunders":
		blundersHandler(w, r)
	default:
		http.NotFound(w, r)
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	router(w, r)
}
