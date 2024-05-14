package main

import (
	"log"

	"github.com/kirkalyn13/xyz-books-app/server/cmd/routes"
	"github.com/kirkalyn13/xyz-books-app/server/pkg/db"
)

func main() {
	log.Println("Starting XYZ Books Server...")

	err := db.LoadDatabase()

	if err != nil {
		log.Fatalf("Error when loading database: %v", err)
	}

	r := routes.Router()

	err = r.Run()

	if err != nil {
		log.Fatalf("Error when running server: %v", err)
	}
}
