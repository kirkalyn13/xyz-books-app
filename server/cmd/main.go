package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kirkalyn13/xyz-books-app/cmd/routes"
)

func main() {
	log.Println("Starting XYZ Books Server...")

	r := gin.Default()
	routes.RegisterRoutes(r)

	r.Run()
}
