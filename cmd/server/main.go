package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/chippolot/haunted-limo/api"
	blundersAPI "github.com/chippolot/haunted-limo/api/blunders"
	blunders "github.com/chippolot/haunted-limo/api/blunders/_internal"
)

var dataProvider *blunders.SQLDataProvider

func main() {
	// Resolve DB connection string
	dsn := os.Getenv("DSN")
	if dsn == "" {
		panic("DSN not found in environment variables")
	}

	dataProvider = blunders.MakeSQLDataProvider(dsn)
	defer dataProvider.Close()

	http.Handle("/", http.HandlerFunc(api.Index))
	http.Handle("/blunders", http.HandlerFunc(blundersAPI.Blunders))
	http.Handle("/blunders/api/cron", http.HandlerFunc(blundersAPI.Cron))

	port := 8080
	fmt.Printf("Server is running on http://localhost:%v\n", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
