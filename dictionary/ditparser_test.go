package dictionary

import (
	"slices"
	"testing"
)

func TestParsingDictionary(t *testing.T) {

	got, err := ParseJSON("testdict.json")
	want := []string{"Word 1", "Word 2"}

	if err != nil {
		t.Errorf("Couldn't read file: %v", err)
	}

	if !slices.Equal(got, want) {
		t.Errorf("ParseJSON(\"dictionary/testdict.json\") = %s; want %s", got, want)
	}
}
