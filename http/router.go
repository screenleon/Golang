package http

import (
	"net/http"
	"server/controllers"

	"github.com/julienschmidt/httprouter"
)

func Router() (handler http.Handler) {
	router := httprouter.New()
	router.GET("/health", controllers.Health)

	return router
}
