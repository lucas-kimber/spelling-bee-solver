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

func (dict wordMap) AddWord(newWord string) {

	dict.store[makeKey(newWord)] = append(dict.store[makeKey(newWord)], newWord)
}
