package dictionary

import (
	"encoding/json"
	"os"
)

// type jsonStructure struct {
// 	field []string
// }

// func ParseJSON(path string) ([]string, error) {

// 	parsedWords := []string{}

// 	file, err := os.Open(path)

// 	if err != nil {
// 		return nil, err
// 	}

// 	defer file.Close()

// 	decoder := json.NewDecoder(file)

// 	for decoder.More() {

// 		var entry jsonStructure

// 		if err := decoder.Decode(&entry); err != nil {
// 			return nil, err
// 		}

// 		parsedWords = append(parsedWords, entry.field[1])
// 	}

// 	return parsedWords, nil
// }

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
