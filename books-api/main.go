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
	{ID: "1", Title: "1984", Author: "George Orwell", Quantity: 3},
	{ID: "2", Title: "To Kill a Mockingbird", Author: "Harper Lee", Quantity: 5},
	{ID: "3", Title: "The Catcher in the Rye", Author: "J.D. Salinger", Quantity: 4},
	{ID: "4", Title: "The Hobbit", Author: "J.R.R. Tolkien", Quantity: 6},
	{ID: "5", Title: "Pride and Prejudice", Author: "Jane Austen", Quantity: 2},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// bookById is a handler that retrieves a book by its ID provided as a URL parameter.
func bookById(c *gin.Context) {
	// Extracts the "id" parameter from the URL.
	var id = c.Param("id")
	// id := c.Param("id")

	// Calls the getBookById function to fetch the book with the given ID.
	book, err := getBookById(id) //The `:=` operator in Go is used for short variable declaration and assignment.

	// Checks if an error occurred while fetching the book.
	if err != nil {
		// Returns a JSON response with HTTP status 404 (Not Found) and an error message.
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found!"})
		return
	}

	// Returns the found book as a JSON response with HTTP status 200 (OK).
	c.IndentedJSON(http.StatusOK, book)
}

func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query param"})
		return
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book no found!"})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available!"})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query param"})
		return
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book no found!"})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}

func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
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
	router.GET("/books/:id", bookById)
	router.POST("/books", createBook)
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/return", returnBook)
	router.Run("localhost:8080")
}
