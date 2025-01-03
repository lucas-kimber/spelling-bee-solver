package solver

import (
	"sort"
	"strings"
)

type solver struct {
	dictionary wordMap
}

func NewSolver() *solver {
	return &solver{*NewWordMap()}
}

func (s solver) ParseWords(words []string) {

	for _, w := range words {
		s.dictionary.AddWord(w)
	}
}

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

func insertInOrder(c string, str string) string {

	str += c
	sortedStr := strings.Split(str, "")

	sort.Strings(sortedStr)

	return strings.Join(sortedStr, "")
}

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
