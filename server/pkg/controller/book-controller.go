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
	searchQuery := c.Query("q")

	results, err := service.GetBooks(searchQuery)

	if err != nil {
		log.Println(err)
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found."})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

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

	result, err := service.GetBookByISBN13(isbn13)

	if err != nil {
		log.Println(err)
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found."})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"book": result,
	})
}

// GetBookByID is the controller to fetch a Book based on a given ID
func GetBookByID(c *gin.Context) {
	id := c.Param("id")

	result, err := service.GetBookByID(id)

	if err != nil {
		log.Println(err)
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found."})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"book": result,
	})
}

// AddBook is the controller to add a new Book entity
func AddBook(c *gin.Context) {
	var book model.Book
	err := c.Bind(&book)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result, err := service.AddBook(book)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"book": result})
}

// EditBook is the controller to edit a Book entity
func EditBook(c *gin.Context) {
	var book model.Book
	id := c.Param("id")

	err := c.Bind(&book)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result, err := service.EditBook(book, id)

	if err != nil {
		log.Println(err)
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found."})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"book": result,
	})

}

// DeleteBook is the controller to delete a Book entity
func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	err := service.DeleteBook(id)

	if err != nil {
		log.Println(err)
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found."})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
