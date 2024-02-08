package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {

	db, err := sql.Open("sqlite3", "./api.db")
	if err != nil {
		panic("Database could not connect: " + err.Error())
	}

	DB = db
	DB.Exec("PRAGMA foreign_keys = ON;")
	err = createTables()
	if err != nil {
		panic("Database could not connect: " + err.Error())
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxOpenConns(5)
	fmt.Println("Tables created successfully!")
}

func createTables() error {
	createBooksTable := `
        CREATE TABLE IF NOT EXISTS books (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL UNIQUE,
            price INTEGER NOT NULL,
            publishDate DATETIME NOT NULL,
			author_id INTEGER,
			FOREIGN KEY(author_id) REFERENCES authors(id)
        )
    `

	createAuthorsTable := `
	CREATE TABLE IF NOT EXISTS authors (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
`
	_, err := DB.Exec(createBooksTable)
	if err != nil {
		return err
	}
	_, err = DB.Exec(createAuthorsTable)
	return err
}
