package datetimeserver

import (
	"errors"
	"io"
	"net/http"
	"time"
)

var (
	ErrInvalidEndpoint = errors.New("invalid endpoint")
	ErrInvalidMethod   = errors.New("invalid method")
	ErrInvalidFormat   = errors.New("invalid format")
)

func GetDateAndTime(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, time.Now().Format("2006-01-02 15:04:05"))
}
