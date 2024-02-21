package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chippolot/jokegen"
	_ "github.com/go-sql-driver/mysql"
	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:  "blunderbuddy",
		Usage: "generate a comical story of a misundertanding!",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "themesPath",
				Aliases:  []string{"t"},
				Value:    "",
				Usage:    "Path to theme strings",
				Required: false,
			},
			&cli.StringFlag{
				Name:     "stylesPath",
				Aliases:  []string{"s"},
				Value:    "",
				Usage:    "Path to style strings",
				Required: false,
			},
			&cli.StringFlag{
				Name:     "modifiersPath",
				Aliases:  []string{"m"},
				Value:    "",
				Usage:    "Path to modifier strings",
				Required: false,
			},
			&cli.StringFlag{
				Name:     "storyType",
				Aliases:  []string{"t"},
				Value:    "",
				Usage:    "Path to modifier strings",
				Required: false,
			},
		},
		Action: func(ctx *cli.Context) error {
			storyTypeString := ctx.String("storyType")
			storyType, err := jokegen.ParseStoryType(storyTypeString)
			if err != nil {
				return err
			}

			themesPath := ctx.String("themesPath")
			stylesPath := ctx.String("stylesPath")
			modifiersPath := ctx.String("modifiersPath")

			return prepareDb(storyType, themesPath, stylesPath, modifiersPath)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func prepareDb(storyType jokegen.StoryType, themesPath, stylesPath, modifiersPath string) error {
	dsn := os.Getenv("DSN")
	if dsn == "" {
		log.Fatal("DSN not found in environment variables")
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	if themesPath != "" {
		if err := insertStringsFromFile(db, -1, "Themes", "Theme", themesPath); err != nil {
			return err
		}
	}
	if stylesPath != "" {
		if err := insertStringsFromFile(db, storyType, "Styles", "Style", stylesPath); err != nil {
			return err
		}
	}
	if modifiersPath != "" {
		if err := insertStringsFromFile(db, storyType, "Modifiers", "Modifier", modifiersPath); err != nil {
			return err
		}
	}

	fmt.Println("Finished preparing database.")
	return nil
}

func insertStringsFromFile(db *sql.DB, storyType jokegen.StoryType, table, column, filePath string) error {
	file, err := jokegen.ResourcesFS.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	sliceSize := 200
	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		if len(lines) == sliceSize {
			insertStrings(db, storyType, table, column, lines)
			lines = []string{}
		}
	}
	insertStrings(db, storyType, table, column, lines)

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func insertStrings(db *sql.DB, storyType jokegen.StoryType, table, column string, data []string) error {
	if len(data) == 0 {
		return nil
	}

	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("Failed to start transaction: %v", err)
	}
	defer tx.Rollback() // The rollback will be ignored if the transaction has been committed

	// Prepare the statement for inserting data, including a timestamp
	maybeStoryTypeColumn, maybeStoryTypeArg := "", ""
	if storyType == -1 {
		maybeStoryTypeColumn = ", StoryType"
		maybeStoryTypeArg = ", ?"
	}
	stmt, err := tx.Prepare(fmt.Sprintf("INSERT INTO %s (%s, Timestamp%s) VALUES (?, ?%s)", table, column, maybeStoryTypeColumn, maybeStoryTypeArg))
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Insert all lines
	for _, entry := range data {
		line := entry
		timestamp := time.Now().UTC()

		_, err := stmt.Exec(line, timestamp)
		if err != nil {
			return err
		}
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		log.Fatalf("Failed to commit transaction: %v", err)
	}

	fmt.Printf("Successfully inserted %v entries into table %s.\n", len(data), table)
	return nil
}
