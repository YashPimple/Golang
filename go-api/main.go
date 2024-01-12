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
	{ID: "1", Title: "In seacrh of lost time", Author: "Marcel Proust", Quantity: 3},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 4},
	{ID: "3", Title: "War and Piece", Author: "Leo Tolstoy", Quantity: 6},
}

var Greetings = "Hello Sir Welcome to a Tutorial of Go-API using gins"

func getHomepage(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Greetings)
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book no not found"})
		return
	}

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

func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok { // ok == false
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter"})
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book no not found"})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available"})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)

}

func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter"})
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}
func main() {

	r := gin.Default()
	r.GET("/", getHomepage)
	r.GET("/books", getBooks)
	r.GET("/books/:id", bookById)
	r.POST("/books", createBook)
	r.PATCH("/checkout", checkoutBook)
	r.PATCH("/return", returnBook)
	r.Run("localhost:8080")
}
