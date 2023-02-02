package dictionary_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/haydenccarroll/definition-api/cmd/definition-api/structs"
)

func GetDefinitions(word string) ([]string, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.dictionaryapi.dev/api/v2/entries/en/%s", word))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data structs.DictionaryResponse

	if err = json.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
		return nil, err
	}

	var definitions []string

	// get first definition for every meaning
	for _, meaning := range data[0].Meanings {
		definitions = append(definitions, meaning.Definitions[0].Definition)
	}

	return definitions, nil
}
