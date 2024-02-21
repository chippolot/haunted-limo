package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/chippolot/haunted-limo/api"
	common "github.com/chippolot/haunted-limo/api/_pkg"
)

var dataProvider *common.SQLDataProvider

func main() {
	// Resolve DB connection string
	dsn := os.Getenv("DSN")
	if dsn == "" {
		panic("DSN not found in environment variables")
	}

	dataProvider = common.MakeSQLDataProvider(dsn)
	defer dataProvider.Close()

	http.Handle("/", http.HandlerFunc(api.Index))
	http.Handle("/blunders", http.HandlerFunc(api.Blunders))
	http.Handle("/whammies", http.HandlerFunc(api.Whammies))
	http.Handle("/cron", http.HandlerFunc(api.Cron))

	port := 8080
	fmt.Printf("Server is running on http://localhost:%v\n", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
