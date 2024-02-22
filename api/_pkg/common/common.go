package common

import (
	"io"
	"os"
)

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

func GetDataFilePath(relativePath string) string {
	baseDataDir := os.Getenv("BASE_DATA_DIR") + "data/"
	return baseDataDir + relativePath
}

func LoadFileBytes(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
