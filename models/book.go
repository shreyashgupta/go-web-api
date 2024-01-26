package models

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	Id          string 
	Name        string `binding:"required"`
	Author      string `binding:"required"`
	Price       int `binding:"required"`
	PublishDate time.Time
}

var Books = []Book{}

func AddBook(book Book) *Book{
	newBook := createBook(book.Name, book.Author, book.Price)
	Books = append(Books, *newBook)
	return newBook
}

func GetAllBooks ()[]Book{
	return Books
}

func createBook(name string, author string, price int) *Book{
	book:= Book{
		Name: name,
		Author: author,
		Price: price,
		PublishDate: time.Now(),
		Id: uuid.New().String(),
	}

	return &book
}