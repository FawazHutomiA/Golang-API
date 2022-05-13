package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type book struct{
	ID		string `json:"id"`
	Title	string `json: "title"`
	Author	string `json: "author"`
	Quantity int   `json: "quantity"`
}

var books = []book{
	{ID: "1", Title: "Mulia dengan Manhaj Salaf", Author: "Fawaz", Quantity: 20},
	{ID: "2", Title: "Tauhid", Author: "Hutomi", Quantity: 30},
	{ID: "3", Title: "Adab Sebelum Amal", Author: "Abdurahman", Quantity: 40},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func postBooks(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func getBookById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range books {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookById)
	router.POST("/books", postBooks)

	router.Run("localhost:8080")
}