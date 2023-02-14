package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/haydenccarroll/definition_api/cmd/definition_api/anki_client"
	"github.com/haydenccarroll/definition_api/cmd/definition_api/dictionary_api"
)

const (
	deckName = "Obscure English Words"
)

func LookupWord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	word := vars["word"]

	definitions, err := dictionary_api.GetDefinitions(word)
	if err != nil {
		log.Println("Error getting definition", err)
	}

	if err = anki_client.AddCard(word, definitions, deckName); err != nil {
		log.Println("Error storing definition", err)
	}

	w.WriteHeader(http.StatusOK)
	resp, err := json.Marshal(definitions)
	if err != nil {
		log.Println("Could not marshal definitions to send in response body", err)
	}

	if _, err = w.Write(resp); err != nil {
		log.Println("Could not write response", err)
	}
}
