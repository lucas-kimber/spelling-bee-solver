package solver

import (
	"slices"
	"testing"
)

func TestInsertInOrder(t *testing.T) {
	got := insertInOrder("a", "bc")
	want := "abc"

	if got != want {
		t.Errorf("insertInOrder(\"a\", \"bc\") = %s; want %s", got, want)
	}
}

func TestFindSubsets(t *testing.T) {
	got := findSubsets("abc")
	want := []string{"a", "b", "ab", "c", "ac", "bc", "abc"}

	if !slices.Equal(got, want) {
		t.Errorf("findSubsets(\"abc\") = %s; want %s", got, want)
	}
}

func TestSolve(t *testing.T) {

	s := NewSolver()
	s.ParseWords([]string{"abcd", "bcad", "bacd", "xyzd"})

	got := s.Solve("a", "bcdefg")
	want := []string{"abcd", "bcad", "bacd"}

	if !slices.Equal(got, want) {
		t.Errorf("ParseWords([]string{\"abcd\", \"bcad\", \"bacd\", \"xyzd\"}) then Solve(\"a\", \"abcdefg\") = %s; want %s", got, want)
	}
}

func TestSolveRejectsWithoutCentreLetter(t *testing.T) {

	s := NewSolver()
	s.ParseWords([]string{"bcde", "cbde", "abcd"})

	got := s.Solve("a", "bcdefg")
	want := []string{"abcd"}

	if !slices.Equal(got, want) {
		t.Errorf("ParseWords([]string{\"bcde\", \"cbde\", \"abcd\"}) then Solve(\"a\", \"abcd\") = %s; want %s", got, want)
	}
}

func TestSolveSubsets(t *testing.T) {

	s := NewSolver()
	s.ParseWords([]string{"abcc", "abccdd", "abba", "defg"})

	got := s.Solve("a", "bcdefg")
	want := []string{"abba", "abcc", "abccdd"}

	if !slices.Equal(got, want) {
		t.Errorf("ParseWords([]string{\"abcc\", \"abccdd\", \"aaaa\"}) then Solve(\"a\", \"abcdef\") = %s; want %s", got, want)
	}
}
