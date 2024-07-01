package main

import (
	"log"
	"net/http"

	datetimeserver "github.com/codescalersinternships/datetime-server-eyadhussein/pkg"
)

func main() {
	http.HandleFunc("/datetime", datetimeserver.GetDateAndTime)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("failed to start the server")
	}
}
