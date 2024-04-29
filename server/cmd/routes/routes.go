package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kirkalyn13/xyz-books-app/pkg/controller"
)

// RegisterRoutes registers the controllers defined
func RegisterRoutes(r *gin.Engine) {
	// GET Requests
	r.GET("/book", controller.GetBook)
}
