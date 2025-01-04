package dictionary

import (
	"encoding/json"
	"os"
)

// ParseJSON takes a path to a dictionary in JSON format and parses it into a slice of strings.
// It assumes the field name for the words in the dictionary to be "word".
func ParseJSON(path string) ([]string, error) {
	parsedWords := []string{}

	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)

	_, err = decoder.Token()

	if err != nil {
		return nil, err
	}

	for decoder.More() {

		var entry struct {
			Word string `json:"word"`
		}

		if err := decoder.Decode(&entry); err != nil {
			return nil, err
		}

		parsedWords = append(parsedWords, entry.Word)
	}

	return parsedWords, nil
}
