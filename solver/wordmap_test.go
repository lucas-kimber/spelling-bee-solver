package solver

import "testing"

func TestMakeKeyCreatesSet(t *testing.T) {
	got := makeKey("aabbcc")
	want := "abc"

	if got != want {
		t.Errorf("makeKey(\"aaccbb\") = %s; want %s", got, want)
	}
}

func TestMakeKeySorts(t *testing.T) {
	got := makeKey("aaccbb")
	want := "abc"

	if got != want {
		t.Errorf("makeKey(\"aaccbb\") = %s; want %s", got, want)
	}
}

func TestMakeKeyEquality(t *testing.T) {
	got := makeKey("aaccbb") == makeKey("bbbbccccaa")
	want := true

	if got != want {
		t.Errorf("makeKey(\"aaccbb\") == makeKey(\"bbbbccccaa\") = %b; want %b", got, want)
	}
}
