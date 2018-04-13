package thesaurus

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BigHuge struct {
	APIKey string
}

type synonyms struct {
	Noun *words `json:"noun"`
	Verb *words `json:"verb"`
}

type words struct {
	Syn []string `json:"syn"`
}

func (b *BigHuge) Synonyms(term string) ([]string, error) {
	var result []string
	response, err := http.Get("http://words.bighugelabs.com/api/2/" + b.APIKey + "/" + term + "/json")
	defer response.Body.Close()
	if err != nil {
		return result, fmt.Errorf("bighuge: failed to search synonyms of %q: %v", term, err)
	}

	var data synonyms
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return result, err
	}
	if data.Noun != nil && data.Noun.Syn != nil {
		result = append(result, data.Noun.Syn...)
	}
	if data.Verb != nil && data.Verb.Syn != nil {
		result = append(result, data.Verb.Syn...)
	}
	return result, nil
}
