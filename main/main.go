package main

import (
	"log"
	"net/http"
	"os"
	"server/services"

	"github.com/julienschmidt/httprouter"
)

func main() {
	port := os.Getenv("PORT")
	router := httprouter.New()
	router.GET("/health", services.Health)
	log.Fatal(http.ListenAndServe(port, router))
}
