package solver

import (
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

func findSubsets(letterSet string) []string {

	arr := strings.Split(letterSet, "")
	sol := []string{}

	for y := 2; y < len(arr); y++ {
		for x := 0; x+y < len(arr); x++ {
			subStr := arr[x : x+y+1]
			sol = append(sol, strings.Join(subStr, ""))
		}
	}

	return sol
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
		solutionSet = append(solutionSet, s.dictionary.Lookup(i)...)
	}

	return solutionSet
}
