// Package datetime provides a handler to return the current date and time.
package datetime

import (
	"io"
	"net/http"
	"time"
)

// GetDateAndTime takes in a response writer and a request then uses the
// response writer to write the current date and time in the format
// "2006-01-02 15:04:05" to the response writer.
//
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
