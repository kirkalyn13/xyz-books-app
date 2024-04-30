package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kirkalyn13/xyz-books-app/pkg/model"
	"github.com/kirkalyn13/xyz-books-app/pkg/service"
	"github.com/kirkalyn13/xyz-books-app/pkg/util"
)

// GetBooks is the controller to fetch the list of Books
func GetBooks(c *gin.Context) {
	results := service.GetBooks()

	c.JSON(http.StatusOK, gin.H{
		"books": results,
	})
}

// GetBooksByISBN13 is the controller to fetch a Book based on a given ISBN13
func GetBookByISBN13(c *gin.Context) {
	isbn13 := c.Param("isbn13")

	if !util.IsValidISBN13(isbn13) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ISBN13."})
		return
	}

	result := service.GetBookByISBN13(isbn13)

	c.JSON(http.StatusOK, gin.H{
		"book": result,
	})
}

// AddBook is the controller to add a new Book entity
func AddBook(c *gin.Context) {
	var book model.Book
	err := c.Bind(&book)

	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	result, err := service.AddBook(book)

	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"book": result})
}

// EditBook is the controller to edit a Book entity
func EditBook(c *gin.Context) {
	var book model.Book
	isbn13 := c.Param("isbn13")

	if !util.IsValidISBN13(isbn13) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ISBN13."})
		return
	}

	err := c.Bind(&book)

	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	result := service.EditBook(book, isbn13)

	c.JSON(http.StatusOK, gin.H{
		"book": result,
	})

}

// DeleteBook is the controller to delete a Book entity
func DeleteBook(c *gin.Context) {
	isbn13 := c.Param("isbn13")

	if !util.IsValidISBN13(isbn13) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ISBN13."})
		return
	}

	service.DeleteBook(isbn13)

	c.Status(http.StatusNoContent)
}
