package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/chippolot/haunted-limo/api"
	"github.com/chippolot/haunted-limo/api/_pkg/blunders"
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

	port := 8080
	fmt.Printf("Server is running on http://localhost:%v\n", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), http.HandlerFunc(api.Handler))
}
