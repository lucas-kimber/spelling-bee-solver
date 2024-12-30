package solver

import (
	"sort"
	"strings"
)

type wordMap struct {
	store map[string][]string
}

func NewWordMap() *wordMap {
	return &wordMap{map[string][]string{}}
}

func makeKey(word string) string {

	arr := strings.Split(word, "")
	setMap := make(map[string]bool)

	for _, i := range arr {
		setMap[i] = true
	}

	setSlice := []string{}

	for k := range setMap {
		setSlice = append(setSlice, k)
	}

	sort.Strings(setSlice)

	return strings.Join(setSlice, "")
}

func (dict wordMap) AddWord(newWord string) {

	dict.store[makeKey(newWord)] = append(dict.store[makeKey(newWord)], newWord)
}
