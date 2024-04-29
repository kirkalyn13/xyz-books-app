package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("STARTING XYZ Books Server...")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
