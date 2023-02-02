package server

import (
	"github.com/gorilla/mux"
	"github.com/haydenccarroll/definition-api/cmd/definition-api/handlers"
)

func getRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/{word}", handlers.LookupWord).Methods("GET")

	return router
}
