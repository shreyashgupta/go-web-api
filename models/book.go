package models

import (
	"time"

	"example.com/go_api/db"
)

type Book struct {
	Id          int64
	Name        string `binding:"required"`
	AuthorId    int64  `binding:"required"`
	Price       int    `binding:"required"`
	PublishDate time.Time
}

func AddBook(book *Book) (int64, error) {
	query := `INSERT INTO books(name, price, publishDate, author_id)
	 VALUES (?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(book.Name, book.Price, book.PublishDate, book.AuthorId)
	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return id, nil
}

func GetAllBooks() ([]Book, error) {
	books := []Book{}
	getAllBooks := `
	SELECT * FROM books;
	`
	rows, err := db.DB.Query(getAllBooks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.Id, &book.Name, &book.Price, &book.PublishDate, &book.AuthorId)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, err
}

func GetBookWithId(id int64) (*Book, error) {
	query := `SELECT * from books where id = (?)`
	rows := db.DB.QueryRow(query, id)
	book := Book{}

	err := rows.Scan(&book.Id, &book.Name, &book.Price, &book.PublishDate, &book.AuthorId)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func UpdateBook(book *Book) error {
	query := `
	UPDATE books
	SET name = ?, author_id = ?, price = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(book.Name, book.Price, book.Id, book.AuthorId)
	if err != nil {
		return err
	}
	return nil
}

func DeleteBook(id int64) error {
	query := `
	DELETE FROM books
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
