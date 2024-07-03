package datetimegin

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	validEndpoint = "/datetime"
	validMethod   = http.MethodGet
	successCode   = http.StatusOK
	serverUrl     = "http://localhost:8080"

	invalidEndpoint            = "/invalid"
	invalidMethod              = "POST"
	methodNotAllowedStatusCode = 405
	notFoundStatusCode         = 404
	internalServerErrorCode    = 500

	validDateTimeFormat   = "2006-01-02 15:04:05"
	invalidDataTimeFormat = "2006-January-02"
)

func TestGetDateTime(t *testing.T) {

	router := gin.Default()
	router.HandleMethodNotAllowed = true

	router.GET("/datetime", GetDateAndTime)

	var getDateTimeStatusCodeTests = []struct {
		testName   string
		endpoint   string
		method     string
		statusCode int
	}{
		{"test get date and time with valid endpoint and method", validEndpoint, validMethod, successCode},
		{"test get date and time with invalid endpoint", invalidEndpoint, validMethod, notFoundStatusCode},
		{"test get date and time with invalid method", validEndpoint, invalidMethod, methodNotAllowedStatusCode},
		{"test get date and time with invalid date time format", validEndpoint, validMethod, successCode},
	}

	for _, tt := range getDateTimeStatusCodeTests {
		t.Run(tt.testName, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, serverUrl+tt.endpoint, nil)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			res := w.Result()

			if res.StatusCode == internalServerErrorCode {
				t.Error("internal server error failed to get date and time")
			}

			if res.StatusCode != tt.statusCode {
				t.Errorf("expected status code %d but got %d", tt.statusCode, res.StatusCode)
			}
		})
	}

	t.Run("test get date and time with valid date time format", func(t *testing.T) {
		req := httptest.NewRequest(validMethod, serverUrl+validEndpoint, nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		res := w.Result()

		if res.StatusCode == internalServerErrorCode {
			t.Error("internal server error failed to get date and time")
		}

		if res.StatusCode != successCode {
			t.Errorf("expected status code %d but got %d", successCode, res.StatusCode)
		}

		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil {
			t.Errorf("failed to read response body")
		}

		dataNoQoutes := string(data)[1 : len(data)-1]
		if dataNoQoutes != time.Now().Format(validDateTimeFormat) {
			t.Errorf("expected date time format %s but got %s", time.Now().Format(validDateTimeFormat), dataNoQoutes)
		}
	})
}
