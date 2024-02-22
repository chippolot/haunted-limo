package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/chippolot/haunted-limo/api"
	"github.com/chippolot/haunted-limo/api/_pkg/data"
	"github.com/joho/godotenv"
)

func cronWithAuthorization(w http.ResponseWriter, r *http.Request) {
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("CRON_SECRET")))
	api.Cron(w, r)
}

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// Resolve DB connection string
	dsn := os.Getenv("DSN")
	if dsn == "" {
		panic("DSN not found in environment variables")
	}

	http.Handle("/", http.HandlerFunc(api.Index))
	http.Handle("/api/cron", http.HandlerFunc(cronWithAuthorization))

	storyDataList, err := data.LoadStoryData()
	if err != nil {
		panic(err)
	}
	for _, storyData := range storyDataList {
		storyType := storyData.StoryType
		http.Handle("/"+storyData.Key, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			query := r.URL.Query()
			query.Add("storyType", storyType)
			r.URL.RawQuery = query.Encode()
			api.Story(w, r)
		}))
	}

	port := 8080
	fmt.Printf("Server is running on http://localhost:%v\n", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
