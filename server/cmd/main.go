package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kirkalyn13/xyz-books-app/server/cmd/routes"
	"github.com/kirkalyn13/xyz-books-app/server/pkg/db"
)

func main() {
	log.Println("Starting XYZ Books Server...")

	err := db.LoadDatabase()

	if err != nil {
		log.Fatalf("Error when loading database: %v", err)
	}

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:5173", "http://localhost:4173"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:  []string{"Origin", "Content-Type"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}))

	routes.RegisterRoutes(r)

	err = r.Run()

	if err != nil {
		log.Fatalf("Error when running server: %v", err)
	}
}
