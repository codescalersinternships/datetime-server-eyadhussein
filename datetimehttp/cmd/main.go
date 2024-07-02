// This file contains the main function that starts the server and listens on port 8080.
package main

import (
	"log"
	"net/http"

	datetimehttp "github.com/codescalersinternships/datetime-server-eyadhussein/datetimehttp/pkg"
)

func main() {
	http.HandleFunc("/datetime", datetimehttp.GetDateAndTime)

	log.Println("Starting server on port 8080...")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatalf("failed to start the server %v", err)
	}
}
