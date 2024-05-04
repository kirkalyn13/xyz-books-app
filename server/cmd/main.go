package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kirkalyn13/xyz-books-app/cmd/routes"
	"github.com/kirkalyn13/xyz-books-app/pkg/db"
)

func main() {
	log.Println("Starting XYZ Books Server...")

	db.LoadDatabase()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:5173"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:  []string{"Origin"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}))

	routes.RegisterRoutes(r)

	r.Run()
}
