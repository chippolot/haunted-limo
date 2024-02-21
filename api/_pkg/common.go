package common

import (
	"html/template"
	"os"
)

type StoryModel struct {
	Title              string
	Story              string
	BackgroundColor    string
	LogoFontLink       template.URL
	LogoFontFamilyName string
	LogoFontStyle      string
	LogoFontWeight     int
	LogoFontSerif      string
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
