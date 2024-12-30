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
		t.Errorf("makeKey(\"aaccbb\") == makeKey(\"bbbbccccaa\") = %t; want %t", got, want)
	}
}

func TestMakeKeyCapitals(t *testing.T) {
	got := makeKey("Ab") == makeKey("Ba")
	want := true

	if got != want {
		t.Errorf("makeKey(\"Aaa\") == makeKey(\"AAa\") = %t; want %t", got, want)
	}
}
