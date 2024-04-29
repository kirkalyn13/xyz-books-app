package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kirkalyn13/xyz-books-app/cmd/routes"
	"github.com/kirkalyn13/xyz-books-app/pkg/db"
)

func main() {
	log.Println("Starting XYZ Books Server...")

	db.LoadDatabase()

	r := gin.Default()
	routes.RegisterRoutes(r)

	r.Run()
}
