package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"example.com/go_api/db"
	"example.com/go_api/models"
	"github.com/gin-gonic/gin"
)

func getBooks(context *gin.Context) {
	books, err := models.GetAllBooks()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get books"})
	}
	context.JSON(http.StatusOK, gin.H{"books": books})
}

func addBook(context *gin.Context) {
	var book models.Book
	err := context.ShouldBindJSON(&book)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	book.PublishDate = time.Now()
	addedId, err := models.AddBook(&book)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Addition Successful", "AddedId": addedId})
}

func getBook(context *gin.Context) {

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to parse id from headers"})
	}
	book, err := models.GetBookWithId(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Failed to fetch book with id due to %s", err)})
	}
	context.JSON(http.StatusOK, gin.H{"message": book})
}

func main() {

	db.InitDB()
	server := gin.Default()

	server.GET("/books", getBooks)
	server.POST("/add_book", addBook)
	server.GET("/books/:id", getBook)
	server.Run(":3000")
}
