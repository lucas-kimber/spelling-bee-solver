package solver

import (
	"sort"
	"strings"
	"unicode"
)

// WordMap wraps a map structure and stores ordered unique letter sets with corresponding words.
// It is used for grouping words that can be constructed from the same letters.
type WordMap struct {
	store map[string][]string
}

// NewWordMap constructs an empty WordMap.
func NewWordMap() *WordMap {
	return &WordMap{map[string][]string{}}
}

// makeKey takes a word as a string and returns a string consisting of all the unique letters present and in order.
func makeKey(word string) string {

	word = strings.ToLower(word)
	setMap := make(map[rune]bool)

	for _, i := range word {
		setMap[i] = true
	}

	setSlice := []rune{}

	for k := range setMap {

		if unicode.IsLetter(k) {
			setSlice = append(setSlice, k)
		}
	}

	sort.Slice(setSlice, func(i, j int) bool {
		return setSlice[i] < setSlice[j]
	})

	return string(setSlice)
}

// AddWord adds a new word to the WordMap.
// Words added are stored under a key consisting of the set of unique letters in the word in order.
func (wm WordMap) AddWord(newWord string) {

	wm.store[makeKey(newWord)] = append(wm.store[makeKey(newWord)], newWord)
}

// Lookup returns all words in the WordMap that have the same key.
// Practically speaking this means returning all words that can be constructed from the same set of letters.
func (wm WordMap) Lookup(letterSet string) []string {
	return wm.store[letterSet]
}
