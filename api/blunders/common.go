package api

import "os"

func getMySQLConnectionString() string {
	// Resolve DB connection string
	connectionString := os.Getenv("DSN")
	if connectionString == "" {
		panic("DSN not found in environment variables")
	}
	return connectionString
}

func getOpenAIToken() string {
	// Resolve API key
	token := os.Getenv("OPEN_AI_API_KEY")
	if token == "" {
		panic("OpenAI API key not found in environment variables")
	}
	return token
}
