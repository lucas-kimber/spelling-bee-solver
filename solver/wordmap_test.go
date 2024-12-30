package solver

import (
	"slices"
	"testing"
)

func TestMakeKeyCreatesSet(t *testing.T) {
	got := makeKey("aabbcc")
	want := "abc"

	if got != want {
		t.Errorf("makeKey(\"aaccbb\") = %s; want %s", got, want)
	}
}

func TestMakeKeySorts(t *testing.T) {
	got := makeKey("aaccbb")
	want := "abc"

	if got != want {
		t.Errorf("makeKey(\"aaccbb\") = %s; want %s", got, want)
	}
}

func TestMakeKeyEquality(t *testing.T) {
	got := makeKey("aaccbb") == makeKey("bbbbccccaa")
	want := true

	if got != want {
		t.Errorf("makeKey(\"aaccbb\") == makeKey(\"bbbbccccaa\") = %t; want %t", got, want)
	}
}

func TestMakeKeyCapitals(t *testing.T) {
	got := makeKey("Ab") == makeKey("Ba")
	want := true

	if got != want {
		t.Errorf("makeKey(\"Aaa\") == makeKey(\"AAa\") = %t; want %t", got, want)
	}
}

func TestMakeKeySymbolAgnostic(t *testing.T) {
	got := makeKey("There's")
	want := makeKey("Theres")

	if got != want {
		t.Errorf("makeKey(\"There's\") = %s; want %s", got, want)
	}
}

func TestWordmapAddingWords(t *testing.T) {
	wm := NewWordMap()
	wm.AddWord("Abacus")
	wm.AddWord("Bacus")

	got := wm.Lookup("abcsu")
	want := []string{"Abacus", "Bacus"}

	if !slices.Equal(got, want) {
		t.Errorf("AddWord(\"Abacus\") and AddWord(\"Bacus\")then calling Lookup(\"abcsu\") = %s; want %s", got, want)
	}
}
