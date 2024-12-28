package main

import (
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
	{ID: "1", Title: "1984", Author: "George Orwell", Quantity: 3},
	{ID: "2", Title: "To Kill a Mockingbird", Author: "Harper Lee", Quantity: 5},
	{ID: "3", Title: "The Catcher in the Rye", Author: "J.D. Salinger", Quantity: 4},
	{ID: "4", Title: "The Hobbit", Author: "J.R.R. Tolkien", Quantity: 6},
	{ID: "5", Title: "Pride and Prejudice", Author: "Jane Austen", Quantity: 2},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.Run("localhost:8080")
}
