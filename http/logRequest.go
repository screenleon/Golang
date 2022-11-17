package http

import (
	"log"
	"net/http"
	"time"
)

func LogRequest(handler http.Handler) http.Handler {
	logFunc := func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		responseWriter := &ResponseWriter{
			ResponseWriter: w,
			StatusCode:     200,
		}
		handler.ServeHTTP(responseWriter, r)
		log.Printf("%s %s %d %dms\n", r.Method, r.URL, responseWriter.StatusCode, time.Since(startTime).Milliseconds())
	}

	return http.HandlerFunc(logFunc)
}
