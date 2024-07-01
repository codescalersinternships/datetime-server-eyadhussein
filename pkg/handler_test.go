package datetimeserver

import (
	"io"
	"net/http/httptest"
	"testing"
	"time"
)

const (
	validEndpoint   = "/datetime"
	invalidEndpoint = "/invalid"
	validMethod     = "GET"
	invalidMethod   = "POST"

	expectedDateTimeFormat = "2006-01-02 15:04:05"
	invalidDataTimeFormat  = "2006-January-02"
)

func TestGetDateTime(t *testing.T) {
	t.Run("test get date and time with valid endpoint", func(t *testing.T) {
		expected := time.Now().Format(expectedDateTimeFormat)

		req := httptest.NewRequest("GET", "http://localhost:8080/datetime", nil)
		w := httptest.NewRecorder()

		GetDateAndTime(w, req)

		res := w.Result()

		if res.StatusCode != 200 {
			t.Errorf("expected status code 200 but got %d", res.StatusCode)
		}

		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)

		if err != nil {
			t.Errorf("failed to read response body %v", err)
		}

		if string(data) != expected {
			t.Errorf("expected %s but got %s", expected, string(data))
		}
	})
}
