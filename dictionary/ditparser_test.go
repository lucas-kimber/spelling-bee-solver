package dictionary

import (
	"slices"
	"testing"
)

func TestParsingTestDictionary(t *testing.T) {

	got, err := ParseJSON("testdict.json")
	want := []string{"Word 1", "Word 2"}

	if err != nil {
		t.Errorf("Couldn't read file: %v", err)
	}

	if !slices.Equal(got, want) {
		t.Errorf("ParseJSON(\"dictionary.json\") = %s; want %s", got, want)
	}
}

func TestParsingFullDictionary(t *testing.T) {

	_, err := ParseJSON("dictionary.json")

	if err != nil {
		t.Errorf("ParseJSON(\"dictionary.json\") failed with error: %v", err)
	}
}
