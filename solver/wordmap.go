package solver

import (
	"sort"
	"strings"
	"unicode"
)

type wordMap struct {
	store map[string][]string
}

func NewWordMap() *wordMap {
	return &wordMap{map[string][]string{}}
}

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

func (wm wordMap) AddWord(newWord string) {

	wm.store[makeKey(newWord)] = append(wm.store[makeKey(newWord)], newWord)
}

func (wm wordMap) Lookup(letterSet string) []string {
	return wm.store[letterSet]
}
