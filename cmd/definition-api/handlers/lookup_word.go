package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/haydenccarroll/definition-api/cmd/definition-api/anki"
	"github.com/haydenccarroll/definition-api/cmd/definition-api/dictionary_api"
)

func LookupWord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	word := vars["word"]

	definitions, err := dictionary_api.GetDefinitions(word)
	if err != nil {
		log.Println("Error getting definition", err)
	}

	if err = anki.AddCard(word, definitions); err != nil {
		log.Println("Error storing definition", err)
	}

	w.WriteHeader(http.StatusOK)
	resp, err := json.Marshal(definitions)
	if err != nil {
		log.Println("Could not marshal definitions to send in response body", err)
	}

	w.Write(resp)
}
