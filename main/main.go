package main

import (
	"log"
	"net/http"
	"os"
	serverHttp "server/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("app.env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	tlsCertificate := os.Getenv("TLS_CERT")
	tlsKey := os.Getenv("TLS_KEY")
	port := os.Getenv("PORT")

	useTls := false
	if tlsCertificate != "" && tlsKey != "" {
		useTls = true
	}

	if port == "" {
		if useTls {
			port = "443"
		} else {
			port = "80"
		}
	}

	router := serverHttp.Router()

	log.Println("listen on", port, "port")
	if useTls {
		log.Fatal(http.ListenAndServeTLS(":"+port, tlsCertificate, tlsKey, serverHttp.LogRequest(router)))
	} else {
		log.Fatal(http.ListenAndServe(":"+port, serverHttp.LogRequest(router)))
	}
}
