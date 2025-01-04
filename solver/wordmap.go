package solver

import (
	"sort"
	"strings"
	"unicode"
)

// wordMap wraps a map structure and stores ordered unique letter sets with corresponding words.
// It is used for grouping words that can be constructed from the same letters.
type wordMap struct {
	store map[string][]string
}

// NewWordMap constructs an empty wordMap.
func NewWordMap() *wordMap {
	return &wordMap{map[string][]string{}}
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

// AddWord adds a new word to the wordMap.
// Words added are stored under a key consisting of the set of unique letters in the word in order.
func (wm wordMap) AddWord(newWord string) {

	wm.store[makeKey(newWord)] = append(wm.store[makeKey(newWord)], newWord)
}

// Lookup returns all words in the wordMap that have the same key.
// Practically speaking this means returning all words that can be constructed from the same set of letters.
func (wm wordMap) Lookup(letterSet string) []string {
	return wm.store[letterSet]
}
