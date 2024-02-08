package server

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"example.com/go_api/models"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(server *gin.Engine) error {

	server.GET("/books", getBooks)
	server.POST("/add_book", addBook)
	server.GET("/books/:id", getBook)
	server.PUT("/update_book/:id", updateBook)
	server.DELETE("/delete_book/:id", deleteBook)
	server.POST("/signup", addAuthor)
	server.POST("/signin", signIn)

	err := server.Run(":3000")
	return err
}

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

func updateBook(context *gin.Context) {
	var book models.Book
	err := context.ShouldBindJSON(&book)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	_, err = models.GetBookWithId(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "No event with specidied id exists"})
		return
	}

	book.Id = id
	err = models.UpdateBook(&book)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Updation Successful"})
}

func deleteBook(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err = models.DeleteBook(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Deletion Successful"})
}

func addAuthor(context *gin.Context) {
	var user models.Author
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	addedId, err := models.AddAuthor(&user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "SignUp Successful", "AddedId": addedId})
}

func signIn(context *gin.Context) {
	var author models.Author
	err := context.ShouldBindJSON(&author)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err = models.ValidateAuthor(&author)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "SignIn successful"})
}
