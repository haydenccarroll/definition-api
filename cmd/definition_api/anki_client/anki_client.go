package anki_client

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/haydenccarroll/definition_api/cmd/definition_api/structs"
)

const (
	addCardAction = "action"
	ankiVersion   = 1
	modelName     = "this"
	endpoint      = "https://someendpoint.com/ankiconnect/post"
)

func AddCard(word string, definitions []string, deckName string) error {

	var backOfCard string
	for _, def := range definitions {
		backOfCard += def + "\n"
	}

	requestStruct := structs.AddCardRequest{
		Action:  addCardAction,
		Version: ankiVersion,
		Params: structs.Params{
			Note: structs.Note{
				DeckName:  deckName,
				ModelName: modelName,
				Fields: structs.Fields{
					Text:  word,
					Extra: backOfCard,
				},
				Tags:    nil,
				Picture: nil,
			},
		},
	}

	requestBody, err := json.Marshal(requestStruct)
	if err != nil {
		log.Println("")
	}

	// Create a HTTP post request
	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Println("Could not make post request to ankiconnect", err)
		return err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println("Error sending request to ankiconnect", err)
	}

	defer response.Body.Close()

	var responseStruct structs.AddCardResponse

	if err = json.NewDecoder(response.Body).Decode(&responseStruct); err != nil {
		log.Println("Error decoding anki connect response body", err)
	}

	return nil
}
