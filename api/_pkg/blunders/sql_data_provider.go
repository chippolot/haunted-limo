package blunders

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/chippolot/jokegen"
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

func (f *SQLDataProvider) AddStory(story, prompt string, storyType jokegen.StoryType) error {
	now := time.Now().UTC()

	sqlInsert := `INSERT INTO Stories (Story, Prompt, Timestamp) VALUES (?, ?, ?)`
	_, err := f.db.Exec(sqlInsert, story, prompt, now)
	if err != nil {
		return err
	}

	return nil
}

func (f *SQLDataProvider) GetMostRecentStory(storyType jokegen.StoryType) (jokegen.StoryResult, error) {
	var result jokegen.StoryResult

	err := f.db.
		QueryRow(fmt.Sprintf("SELECT Story, Prompt, Timestamp FROM Stories WHERE StoryType = %d ORDER BY Id DESC LIMIT 1", storyType)).
		Scan(&result.Story, &result.Prompt, &result.Timestamp)
	if err != nil && err != sql.ErrNoRows {
		return jokegen.StoryResult{}, err
	}

	return result, nil
}

func (f *SQLDataProvider) GetRandomString(dataType jokegen.StoryDataType, storyType jokegen.StoryType) (string, error) {
	table, column, hasStoryTypeFilter, err := getTableAndColumnName(dataType)
	if err != nil {
		return "", err
	}

	where := ""
	if hasStoryTypeFilter {
		where = fmt.Sprintf("WHERE StoryType = %d", storyType)
	}
	query := fmt.Sprintf("SELECT %s FROM %s %s ORDER BY RAND() LIMIT 1;", column, table, where)

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

func getTableAndColumnName(dataType jokegen.StoryDataType) (string, string, bool, error) {
	switch dataType {
	case jokegen.Themes:
		return "Themes", "Theme", false, nil
	case jokegen.Styles:
		return "Styles", "Style", true, nil
	case jokegen.Modifiers:
		return "Modifiers", "Modifier", true, nil
	}
	return "", "", false, fmt.Errorf("unknown data type %v", dataType)
}
