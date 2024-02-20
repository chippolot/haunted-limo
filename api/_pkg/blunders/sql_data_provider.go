package blunders

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/chippolot/blunder"
	_ "github.com/go-sql-driver/mysql"
)

type SQLDataProvider struct {
	db *sql.DB
}

func MakeSQLDataProvider(connectionString string) *SQLDataProvider {
	var err error

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	dataProvider := &SQLDataProvider{
		db: db,
	}

	return dataProvider
}

func (f *SQLDataProvider) AddStory(story string, prompt string) error {
	now := time.Now().UTC()

	sqlInsert := `INSERT INTO Stories (Story, Prompt, Timestamp) VALUES (?, ?, ?)`
	_, err := f.db.Exec(sqlInsert, story, prompt, now)
	if err != nil {
		return err
	}

	return nil
}

func (f *SQLDataProvider) GetMostRecentStory() (blunder.StoryResult, error) {
	var result blunder.StoryResult

	err := f.db.
		QueryRow("SELECT Story, Prompt, Timestamp FROM Stories ORDER BY Id DESC LIMIT 1").
		Scan(&result.Story, &result.Prompt, &result.Timestamp)
	if err != nil && err != sql.ErrNoRows {
		return blunder.StoryResult{}, err
	}

	return result, nil
}

func (f *SQLDataProvider) GetRandomString(dataType blunder.StoryDataType) (string, error) {
	table, column, err := getTableAndColumnName(dataType)
	if err != nil {
		return "", err
	}

	query := fmt.Sprintf("SELECT %s FROM %s ORDER BY RAND() LIMIT 1;", column, table)

	var str string

	// Execute the query
	err = f.db.QueryRow(query).Scan(&str)
	if err != nil {
		return "", err
	}

	return str, nil
}

func (f *SQLDataProvider) Close() error {
	return f.db.Close()
}

func getTableAndColumnName(dataType blunder.StoryDataType) (string, string, error) {
	switch dataType {
	case blunder.Themes:
		return "Themes", "Theme", nil
	case blunder.Styles:
		return "Styles", "Style", nil
	case blunder.Modifiers:
		return "Modifiers", "Modifier", nil
	}
	return "", "", fmt.Errorf("unknown data type %v", dataType)
}
