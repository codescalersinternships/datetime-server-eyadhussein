// Package datetimegin provides a handler to return the current date and time using Gin.
package datetimegin

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetDateAndTime handles HTTP GET request by reciving a gin context with request
// information and returns the current date and time in the format "2006-01-02 15:04:05".
//
// If the request URL is not "/datetime", it responds with a 404 Not Found status.
// If the request method is not GET, it responds with a 405 Method Not Allowed status.
// If the request is valid, it responds with a 200 OK status and the current date and time.
func GetDateAndTime(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, time.Now().Format("2006-01-02 15:04:05"))
}
