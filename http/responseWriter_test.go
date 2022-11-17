package http

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteHeader(t *testing.T) {
	w := httptest.NewRecorder()
	responseWriter := &ResponseWriter{
		ResponseWriter: w,
		StatusCode:     400,
	}
	responseWriter.WriteHeader(http.StatusOK)
	if status := responseWriter.StatusCode; status != http.StatusOK {
		t.Errorf("Wrong status")
	}
}
