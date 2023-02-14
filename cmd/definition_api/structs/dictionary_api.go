package structs

type DictionaryResponse []entry

type entry struct {
	Word       string     `json:"word"`
	Phonetics  []phonetic `json:"phonetics"`
	Meanings   []meaning  `json:"meanings"`
	License    license    `json:"license"`
	SourceUrls []string   `json:"sourceUrls"`
}

type license struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type meaning struct {
	PartOfSpeech string       `json:"partOfSpeech"`
	Definitions  []definition `json:"definitions"`
	Synonyms     []string     `json:"synonyms"`
	Antonyms     []string     `json:"antonyms"`
}

type definition struct {
	Definition string        `json:"definition"`
	Synonyms   []interface{} `json:"synonyms"`
	Antonyms   []interface{} `json:"antonyms"`
	Example    *string       `json:"example,omitempty"`
}

type phonetic struct {
	Audio     string   `json:"audio"`
	SourceURL *string  `json:"sourceUrl,omitempty"`
	License   *license `json:"license,omitempty"`
	Text      *string  `json:"text,omitempty"`
}
