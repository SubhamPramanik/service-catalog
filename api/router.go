package api

import (
	"kong/api/handler"
	"kong/api/middleware"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	// Apply logging middleware to log all requests
	router.Use(middleware.LoggingMiddleware)

	router.HandleFunc("/services", handler.GetServices).Methods("GET")
	router.HandleFunc("/services/{id}", handler.GetServiceByID).Methods("GET")
	router.HandleFunc("/services/{id}/versions/{versionId}", handler.GetVersionByID).Methods("GET")

	return router
}
