package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kirkalyn13/xyz-books-app/pkg/controller"
)

// RegisterRoutes registers the controllers defined
func RegisterRoutes(r *gin.Engine) {
	// GET Requests
	r.GET("/books", controller.GetBooks)
	r.GET("/books/:isbn13", controller.GetBookByISBN13)

	// POST Requests
	r.POST("/books", controller.AddBook)

	// PUT Requests
	r.PUT("/books/:isbn13", controller.EditBook)

	// DELETE Requests
	r.DELETE("/books/:isbn13", controller.DeleteBook)

}
