package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

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
				Required: true,
			},
			&cli.StringFlag{
				Name:     "stylesPath",
				Aliases:  []string{"s"},
				Value:    "",
				Usage:    "Path to style strings",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "modifiersPath",
				Aliases:  []string{"m"},
				Value:    "",
				Usage:    "Path to modifier strings",
				Required: true,
			},
		},
		Action: func(ctx *cli.Context) error {
			themesPath := ctx.String("themesPath")
			stylesPath := ctx.String("stylesPath")
			modifiersPath := ctx.String("modifiersPath")
			return prepareDb(themesPath, stylesPath, modifiersPath)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func prepareDb(themesPath, stylesPath, modifiersPath string) error {
	dsn := os.Getenv("DSN")
	if dsn == "" {
		log.Fatal("DSN not found in environment variables")
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	if err := insertStringsFromFile(db, "Themes", "Theme", themesPath); err != nil {
		return err
	}
	if err := insertStringsFromFile(db, "Styles", "Style", stylesPath); err != nil {
		return err
	}
	if err := insertStringsFromFile(db, "Modifiers", "Modifier", modifiersPath); err != nil {
		return err
	}

	fmt.Println("Finished preparing database.")
	return nil
}

func insertStringsFromFile(db *sql.DB, table string, column string, filePath string) error {
	file, err := os.Open(filePath)
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
			insertStrings(db, table, column, lines)
			lines = []string{}
		}
	}
	insertStrings(db, table, column, lines)

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func insertStrings(db *sql.DB, table, column string, data []string) error {
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
	stmt, err := tx.Prepare(fmt.Sprintf("INSERT INTO %s (%s, Timestamp) VALUES (?, ?)", table, column))
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
