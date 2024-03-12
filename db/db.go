package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		fmt.Println(err)
		panic("Error connecting to DB")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `
	create table if not exists users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		fmt.Println(err)
		panic("Couldn't create users table")
	}

	createEventsTable := `
	create table if not exists events (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    name TEXT NOT NULL,
	    description TEXT NOT NULL,
	    location TEXT NOT NULL,
	    date_time DATETIME NOT NULL,
	    user_id INTEGER,
	    FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		fmt.Println(err)
		panic("Couldn't create events table")
	}

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY (event_id) REFERENCES events(id),
		FOREIGN KEY (user_id) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createRegistrationsTable)
	if err != nil {
		fmt.Println(err)
		panic("Couldn't create registrations table")
	}
}
