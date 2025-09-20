package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Initialize() error {
	connectionString := os.Getenv("DB_CONNECTION_STRING")
	var err error

	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	fmt.Println("Database connection established")

	err = createTables(context.Background())
	if err != nil {
		return fmt.Errorf("Failed to create tables: %w", err)
	}

	return nil
}

func Close() error {
	fmt.Println("Closing database connection")
	return db.Close()
}
