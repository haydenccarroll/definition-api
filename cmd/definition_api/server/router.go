package server

import (
	"github.com/gorilla/mux"
	"github.com/haydenccarroll/definition_api/cmd/definition_api/handlers"
)

func getRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/{word}", handlers.LookupWord).Methods("GET")

	return router
}
