package solver

import (
	"sort"
	"strings"
	"unicode"
)

// Solver wraps a WordMap and is used to implement functionality for actually solving SpellingBee given a problem set.
type Solver struct {
	dictionary WordMap
}

// NewSolver constructs a Solver with an initialised WordMap
func NewSolver() *Solver {
	return &Solver{*NewWordMap()}
}

// ParseWords takes a slice of strings and adds them to the WordMap if they are longer than three letters.
func (s Solver) ParseWords(words []string) {

	sort.Slice(words, func(i, j int) bool {
		return words[i] < words[j]
	})

	for i, w := range words {

		if i > 0 && w == words[i-1] {
			continue
		}

		foundInvalid := false

		for _, r := range w {

			if !unicode.IsLetter(r) {
				foundInvalid = true
				continue
			}
		}

		if foundInvalid {
			continue
		}

		if len(w) > 3 {

			s.dictionary.AddWord(w)
		}
	}
}

// findSubsets takes a string and generates all unique permutations of its letters.
func findSubsets(s string) []string {

	n := len(s)
	subsets := []string{}

	for i := 1; i < (1 << n); i++ {

		var subset string

		for j := 0; j < n; j++ {

			if (i & (1 << j)) != 0 {

				subset += string(s[j])
			}
		}

		subsets = append(subsets, subset)
	}

	return subsets
}

// insertInOrder takes a string a character, then returns a string with the character added in order.
func insertInOrder(c string, str string) string {

	str += c
	sortedStr := strings.Split(str, "")

	sort.Strings(sortedStr)

	return strings.Join(sortedStr, "")
}

// Solve takes a SpellingBee problem instance and returns a slice containing all valid words.
// The problem set is given as:
// centreLetter, the letter that must be present in all solution words
// inputLetters, the set of letters that can be used to make words
// It then finds all unique permutations of inputLetters, adds the centreLetter, and looks up the words that can be made from these permutations in the WordMap.
// The return is a string slice which is the uninion of all the lookups.
func (s Solver) Solve(centreLetter string, inputLetters string) []string {

	letterSet := ""
	arr := strings.Split(inputLetters, "")

	for _, l := range arr {
		if centreLetter != l {
			letterSet += l
		}
	}

	letterSet = makeKey(letterSet)

	subsets := findSubsets(letterSet)

	solutionSet := []string{}

	for _, i := range subsets {
		solutionSet = append(solutionSet, s.dictionary.Lookup(insertInOrder(centreLetter, i))...)
	}

	sort.Slice(solutionSet, func(i, j int) bool {
		return solutionSet[i] < solutionSet[j]
	})

	return solutionSet
}
