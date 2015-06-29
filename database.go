package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rubenv/sql-migrate"
	"os"
)

var db *sql.DB

func dbconnect() {
	// Connect to database
	db, err := sql.Open("sqlite3", *database)
	if err != nil {
		fmt.Printf("Error connecting to database: %s\n", err)
		os.Exit(3)
	}

	// Ping the database
	err = db.Ping()
	if err != nil {
		fmt.Printf("Error pinging database: %s\n", err)
		os.Exit(4)
	}

	// Setup migrations
	migrations := &migrate.FileMigrationSource{
		Dir: "app/database/migrations",
	}

	// Run migrations
	n, err := migrate.Exec(db, "sqlite3", migrations, migrate.Up)
	if err != nil {
		fmt.Printf("Error running database migrations: %s\n", err)
		os.Exit(5)
	} else {
		if n == 0 {
			fmt.Println("Nothing to migrate")
		} else {
			fmt.Printf("Applied %d migrations\n", n)
		}
	}
}
