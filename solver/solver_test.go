package solver

import (
	"slices"
	"testing"
)

func TestFindSubsets(t *testing.T) {
	got := findSubsets("abcdef")
	want := []string{"abc", "bcd", "cde", "def", "abcd", "bcde", "cdef", "abcde", "bcdef", "abcdef"}

	if !slices.Equal(got, want) {
		t.Errorf("findSubsets(\"abcdef\") = %s; want %s", got, want)
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
	s.ParseWords([]string{"abcc", "abccdd", "aaaa"})

	got := s.Solve("a", "abcdef")
	want := []string{"abcc", "abccdd", "aaaa"}

	if !slices.Equal(got, want) {
		t.Errorf("ParseWords([]string{\"abcc\", \"abccdd\", \"aaaa\"}) then Solve(\"a\", \"abcdef\") = %s; want %s", got, want)
	}
}
