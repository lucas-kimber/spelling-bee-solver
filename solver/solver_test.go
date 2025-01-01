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
