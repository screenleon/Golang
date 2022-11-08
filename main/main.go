package main

import (
	"log"
	"net/http"
	"os"
	"server/controllers"

	"github.com/julienschmidt/httprouter"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	router := httprouter.New()
	router.GET("/health", controllers.Health)

	log.Println("listen on", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
