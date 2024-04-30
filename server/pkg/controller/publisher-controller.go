package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kirkalyn13/xyz-books-app/pkg/model"
	"github.com/kirkalyn13/xyz-books-app/pkg/service"
)

// GetPublishers is the controller to fetch the list of Publishers
func GetPublishers(c *gin.Context) {
	results := service.GetPublishers()

	c.JSON(http.StatusOK, gin.H{
		"publishers": results,
	})
}

// GetPublishersByID is the controller to fetch a Publisher based on a given ID
func GetPublisherByID(c *gin.Context) {
	id := c.Param("id")

	result := service.GetPublisherByID(id)

	c.JSON(http.StatusOK, gin.H{
		"publisher": result,
	})
}

// AddPublisher is the controller to add a new Publisher entity
func AddPublisher(c *gin.Context) {
	var Publisher model.Publisher
	err := c.Bind(&Publisher)

	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	result, err := service.AddPublisher(Publisher)

	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"publisher": result})
}

// EditPublisher is the controller to edit a Publisher entity
func EditPublisher(c *gin.Context) {
	var Publisher model.Publisher
	id := c.Param("id")

	err := c.Bind(&Publisher)

	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	result := service.EditPublisher(Publisher, id)

	c.JSON(http.StatusOK, gin.H{
		"publisher": result,
	})

}

// DeletePublisher is the controller to delete a Publisher entity
func DeletePublisher(c *gin.Context) {
	id := c.Param("id")

	service.DeletePublisher(id)

	c.Status(http.StatusNoContent)
}
