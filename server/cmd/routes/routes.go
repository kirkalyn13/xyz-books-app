package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kirkalyn13/xyz-books-app/pkg/controller"
)

// RegisterRoutes registers the controllers defined
func RegisterRoutes(r *gin.Engine) {
	// GET Requests
	r.GET("/api/v1/books", controller.GetBooks)
	r.GET("/api/v1/books/:isbn13", controller.GetBookByISBN13)
	r.GET("/api/v1/authors", controller.GetAuthors)
	r.GET("/api/v1/authors/:id", controller.GetAuthorByID)
	r.GET("/api/v1/publishers", controller.GetPublishers)
	r.GET("/api/v1/publishers/:id", controller.GetPublisherByID)

	// POST Requests
	r.POST("/api/v1/books", controller.AddBook)
	r.POST("/api/v1/authors", controller.AddAuthor)
	r.POST("/api/v1/publishers", controller.AddPublisher)

	// PUT Requests
	r.PUT("/api/v1/books/:isbn13", controller.EditBook)
	r.PUT("/api/v1/authors/:id", controller.EditAuthor)
	r.PUT("/api/v1/publishers/:id", controller.EditPublisher)

	// DELETE Requests
	r.DELETE("/api/v1/books/:isbn13", controller.DeleteBook)
	r.DELETE("/api/v1/authors/:isbn13", controller.DeleteAuthor)
	r.DELETE("/api/v1/publishers/:isbn13", controller.DeletePublisher)

}
