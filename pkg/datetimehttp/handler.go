// Package datetime provides a handler to return the current date and time using http.
package datetimehttp

import (
	"io"
	"net/http"
	"time"
)

// GetDateAndTime takes in a response writer and a request then uses the
// response writer to write the current date and time in the format
// "2006-01-02 15:04:05" to the response writer.
//
// If the request URL is not "/datetime", it responds with a 404 Not Found status.
// If the request method is not GET, it responds with a 405 Method Not Allowed status.
// If the request is valid, it responds with a 200 OK status and the current date and time.
// If there is an error writing the date and time to the response writer,
// the function writes a 500 internal server error to the response writer.
func GetDateAndTime(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/datetime" {
		http.NotFound(w, r)
		return
	}

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	_, err := io.WriteString(w, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
