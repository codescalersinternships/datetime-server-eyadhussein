package main

import (
	"log"

	"github.com/codescalersinternships/datetime-server-eyadhussein/pkg/datetimegin"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.HandleMethodNotAllowed = true
	router.GET("/datetime", datetimegin.GetDateAndTime)

	log.Println("Starting server on port 8080...")

	err := router.Run("localhost:8080")

	if err != nil {
		log.Fatalf("failed to start the server %v", err)
	}
}
