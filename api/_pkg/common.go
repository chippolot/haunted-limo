package common

import (
	"os"
)

type StoryModel struct {
	Story              string
	BackgroundColor    string
	LogoFontLink       string
	LogoFontFamilyName string
}

func GetMySQLConnectionString() string {
	// Resolve DB connection string
	connectionString := os.Getenv("DSN")
	if connectionString == "" {
		panic("DSN not found in environment variables")
	}
	return connectionString
}

func GetOpenAIToken() string {
	// Resolve API key
	token := os.Getenv("OPEN_AI_API_KEY")
	if token == "" {
		panic("OpenAI API key not found in environment variables")
	}
	return token
}
