package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kirkalyn13/xyz-books-app/pkg/model"
	"github.com/kirkalyn13/xyz-books-app/pkg/service"
)

// GetAuthors is the controller to fetch the list of Authors
func GetAuthors(c *gin.Context) {
	results, err := service.GetAuthors()

	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"authors": results,
	})
}

// GetAuthorsByID is the controller to fetch a Author based on a given ID
func GetAuthorByID(c *gin.Context) {
	id := c.Param("id")

	result, err := service.GetAuthorByID(id)

	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"author": result,
	})
}

// AddAuthor is the controller to add a new Author entity
func AddAuthor(c *gin.Context) {
	var author model.Author
	err := c.Bind(&author)

	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	result, err := service.AddAuthor(author)

	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"author": result})
}

// EditAuthor is the controller to edit a Author entity
func EditAuthor(c *gin.Context) {
	var author model.Author
	id := c.Param("id")

	err := c.Bind(&author)

	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	result, err := service.EditAuthor(author, id)

	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"author": result,
	})

}

// DeleteAuthor is the controller to delete a Author entity
func DeleteAuthor(c *gin.Context) {
	id := c.Param("id")

	err := service.DeleteAuthor(id)

	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.Status(http.StatusNoContent)
}
