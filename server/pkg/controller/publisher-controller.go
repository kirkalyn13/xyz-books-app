package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kirkalyn13/xyz-books-app/server/pkg/model"
	"github.com/kirkalyn13/xyz-books-app/server/pkg/service"
)

// GetPublishers is the controller to fetch the list of Publishers
func GetPublishers(c *gin.Context) {
	searchQuery := c.Query("q")

	results, err := service.GetPublishers(searchQuery)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"publishers": results,
	})
}

// GetPublishersByID is the controller to fetch a Publisher based on a given ID
func GetPublisherByID(c *gin.Context) {
	id := c.Param("id")

	result, err := service.GetPublisherByID(id)

	if err != nil {
		log.Println(err)
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Publisher not found."})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"publisher": result,
	})
}

// AddPublisher is the controller to add a new Publisher entity
func AddPublisher(c *gin.Context) {
	var publisher model.Publisher
	err := c.Bind(&publisher)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result, err := service.AddPublisher(publisher)

	if err != nil {
		log.Println(err)
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Publisher not found."})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"publisher": result})
}

// EditPublisher is the controller to edit a Publisher entity
func EditPublisher(c *gin.Context) {
	var publisher model.Publisher
	id := c.Param("id")

	err := c.Bind(&publisher)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result, err := service.EditPublisher(publisher, id)

	if err != nil {
		log.Println(err)
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Publisher not found."})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"publisher": result,
	})

}

// DeletePublisher is the controller to delete a Publisher entity
func DeletePublisher(c *gin.Context) {
	id := c.Param("id")

	err := service.DeletePublisher(id)

	if err != nil {
		log.Println(err)
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Publisher not found."})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
