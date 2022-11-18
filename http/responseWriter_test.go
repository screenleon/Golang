package http

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteHeader(t *testing.T) {
	w := httptest.NewRecorder()
	wrapResponseWriter := &ResponseWriter{
		ResponseWriter: w,
		StatusCode:     400,
	}
	wrapResponseWriter.WriteHeader(http.StatusOK)
	if status := wrapResponseWriter.StatusCode; status != http.StatusOK {
		t.Errorf("Wrong status")
	}
}
