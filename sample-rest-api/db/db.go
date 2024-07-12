package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "app.db")

	if err != nil {
		panic("Could not connect to DB!!!")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {

	createUsersTable := `
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL
)`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("users table could not be created!!!")
	}

	createEventTable := `
CREATE TABLE IF NOT EXISTS events (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    location TEXT NOT NULL,
    dateTime DATETIME NOT NULL,
    user_id INTEGER,
    FOREIGN KEY(user_id) REFERENCES users(id)
)
`
	_, err = DB.Exec(createEventTable)
	if err != nil {
		panic("event table could not be created!!!")
	}

	createRegistrationTable := `
	CREATE TABLE IF NOT EXISTS registrations (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    event_id INTEGER,
    user_id INTEGER,
    FOREIGN KEY(user_id) REFERENCES users(id),	                                         
    FOREIGN KEY(event_id) REFERENCES events(id)	                                         
)
`
	_, err = DB.Exec(createRegistrationTable)
	if err != nil {
		panic("registrations table could not be created!!!")
	}
}
