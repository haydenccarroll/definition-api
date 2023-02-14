package structs

type AddCardRequest struct {
	Action  string `json:"action"` 
	Version int64  `json:"version"`
	Params  Params `json:"params"` 
}

type Params struct {
	Note Note `json:"note"`
}

type Note struct {
	DeckName  string    `json:"deckName"` 
	ModelName string    `json:"modelName"`
	Fields    Fields    `json:"fields"`   
	Tags      []string  `json:"tags"`     
	Picture   []Picture `json:"picture"`  
}

type Fields struct {
	Text  string `json:"Text"` 
	Extra string `json:"Extra"`
}

type Picture struct {
	URL      string   `json:"url"`
	Filename string   `json:"filename"`
	Fields   []string `json:"fields"`
}

type AddCardResponse struct {
	Result int64       `json:"result"`
	Error  interface{} `json:"error"` 
}

