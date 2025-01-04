package solver

import (
	"sort"
	"strings"
	"unicode"
)

// solver wraps a wordMap and is used to implement functionality for actually solving SpellingBee given a problem set.
type solver struct {
	dictionary wordMap
}

// NewSolver constructs a solver with an initialised wordMap
func NewSolver() *solver {
	return &solver{*NewWordMap()}
}

// ParseWords takes a slice of strings and adds them to the wordMap if they are longer than three letters.
func (s solver) ParseWords(words []string) {

	for _, w := range words {

		for _, r := range w {

			if !unicode.IsLetter(r) {
				continue
			}
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
// It then finds all unique permutations of inputLetters, adds the centreLetter, and looks up the words that can be made from these permutations in the wordMap.
// The return is a string slice which is the uninion of all the lookups.
func (s solver) Solve(centreLetter string, inputLetters string) []string {

	letterSet := ""
	arr := strings.Split(inputLetters, "")

	// Ensure the centre letter isn't in the input letters
	for _, l := range arr {
		if centreLetter != l {
			letterSet += l
		}
	}

	// Ensures it's correctly formed
	letterSet = makeKey(letterSet)

	subsets := findSubsets(letterSet)

	solutionSet := []string{}

	for _, i := range subsets {
		solutionSet = append(solutionSet, s.dictionary.Lookup(insertInOrder(centreLetter, i))...)
	}

	return solutionSet
}
