package main

import (
	"net/http"

	"example.com/go_api/models"
	"github.com/gin-gonic/gin"
)

func getBooks(context *gin.Context){
	books := models.GetAllBooks()
	context.JSON(http.StatusOK, gin.H{"books": books})
}

func addBook(context *gin.Context){
	var book models.Book
	err:= context.ShouldBindJSON(&book)
	if err!=nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	addedBook:=models.AddBook(book)
	context.JSON(http.StatusOK, gin.H{"message": "Addition Successful", "book": *addedBook})
}

func main() {
	server := gin.Default()

	server.GET("/books", getBooks)
	server.POST("/add_book", addBook)
	server.Run(":3000")
}