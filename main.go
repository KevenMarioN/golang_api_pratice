package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "Bom dia Espirito Santo", Author: "Benny Hinn", Quantity: 20},
	{ID: "2", Title: "Bem-vindo Espirito Santo", Author: "Benny Hinn", Quantity: 5},
	{ID: "3", Title: "O falar em línguas", Author: "Luciano Súbira", Quantity: 50},
}

func bookByid(context *gin.Context) {
	id := context.Param("id")
	book, err := getBookById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	context.IndentedJSON(http.StatusOK, book)
}

func getBookById(id string) (*book, error) {
	for index, book := range books {
		if book.ID == id {
			return &books[index], nil
		}
	}
	return nil, errors.New("Book not found!")
}

func getBooks(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, books)
}

func createBook(context *gin.Context) {
	var newBook book

	if err := context.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	context.IndentedJSON(http.StatusCreated, newBook)
}
func checkoutBook(context *gin.Context) {
	id, ok := context.GetQuery("id")
	if !ok {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Query params without id"})
		return
	}

	book, err := getBookById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	if book.Quantity <= 0 {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Without this book"})
		return
	}

	book.Quantity -= 1
	context.IndentedJSON(http.StatusOK, gin.H{"message": "This book available in storage"})
}
func returnBook(context *gin.Context) {
	id, ok := context.GetQuery("id")
	if !ok {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Query params without id"})
		return
	}

	book, err := getBookById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	book.Quantity += 1
	context.IndentedJSON(http.StatusOK, book)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookByid)
	router.POST("/books", createBook)
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/return", returnBook)
	router.Run("localhost:8080")
}
