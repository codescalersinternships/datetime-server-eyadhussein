package datetimeserver

import (
	"io"
	"net/http"
	"time"
)

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
