package http

import (
	"net/http"
)

type ResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func (r *ResponseWriter) WriteHeader(statusCode int) {
	r.StatusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}
