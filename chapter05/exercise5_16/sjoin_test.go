package variadic_test

import (
	"testing"
	"variadic"
)

func TestStringJoinEmptyString(t *testing.T) {
	list := []string{}
	var want string = ""

	got := variadic.StringJoin("", list...)

	if got != want {
		t.Errorf("want: %s, got: %s.", want, got)
	}
}

func TestStringJoinUnpack(t *testing.T) {
	list := []string{"The", "quick", "brown", "fox", "jumped", "over", "the", "lazy", "dog"}
	var want string = "The quick brown fox jumped over the lazy dog"

	got := variadic.StringJoin(" ", list...)

	if got != want {
		t.Errorf("want: %s, got: %s.", want, got)
	}
}

func TestStringJoin(t *testing.T) {
	var want string = "The quick brown fox jumped over the lazy dog"

	got := variadic.StringJoin(" ", "The", "quick", "brown", "fox", "jumped", "over", "the", "lazy", "dog")

	if got != want {
		t.Errorf("want: %s, got: %s.", want, got)
	}
}
