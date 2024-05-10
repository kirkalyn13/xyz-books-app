package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kirkalyn13/xyz-books-app/server/pkg/controller"
)

// RegisterRoutes registers the controllers defined
func RegisterRoutes(r *gin.Engine) {
	// GET Requests
	r.GET("/api/v1/books", controller.GetBooks)
	r.GET("/api/v1/books/:id", controller.GetBookByID)
	r.GET("/api/v1/books/isbn13/:isbn13", controller.GetBookByISBN13)
	r.GET("/api/v1/authors", controller.GetAuthors)
	r.GET("/api/v1/authors/:id", controller.GetAuthorByID)
	r.GET("/api/v1/publishers", controller.GetPublishers)
	r.GET("/api/v1/publishers/:id", controller.GetPublisherByID)

	// POST Requests
	r.POST("/api/v1/books", controller.AddBook)
	r.POST("/api/v1/authors", controller.AddAuthor)
	r.POST("/api/v1/publishers", controller.AddPublisher)

	// PUT Requests
	r.PUT("/api/v1/books/:id", controller.EditBook)
	r.PUT("/api/v1/authors/:id", controller.EditAuthor)
	r.PUT("/api/v1/publishers/:id", controller.EditPublisher)

	// DELETE Requests
	r.DELETE("/api/v1/books/:id", controller.DeleteBook)
	r.DELETE("/api/v1/authors/:id", controller.DeleteAuthor)
	r.DELETE("/api/v1/publishers/:id", controller.DeletePublisher)

}
