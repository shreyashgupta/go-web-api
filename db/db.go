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
            name TEXT NOT NULL,
            author TEXT NOT NULL,
            price INTEGER NOT NULL,
            publishDate DATETIME NOT NULL
        )
    `

	_, err := DB.Exec(createBooksTable)
	return err
}
