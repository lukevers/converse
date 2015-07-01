package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rubenv/sql-migrate"
	"log"
)

var db *sqlx.DB

func dbconnect() {
	// Connect to database
	var err error
	db, err = sqlx.Open("sqlite3", *database)
	if err != nil {
		log.Fatal("Error connecting to database: %s\n", err)
	}

	// Ping the database
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database: %s\n", err)
	}

	// Setup migrations
	migrations := &migrate.FileMigrationSource{
		Dir: "app/database/migrations",
	}

	// Run migrations
	n, err := migrate.Exec(db.DB, "sqlite3", migrations, migrate.Up)
	if err != nil {
		log.Fatal("Error running database migrations: %s\n", err)
	} else {
		if n == 0 {
			log.Println("Nothing to migrate")
		} else {
			log.Println("Applied %d migrations", n)
		}
	}
}
