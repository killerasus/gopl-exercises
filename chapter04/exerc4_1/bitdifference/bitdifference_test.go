package bitdifference_test

import (
	"crypto/sha256"
	"killerasus/exerc4_1/bitdifference"
	"testing"
)

const (
	c1 = "12345"
	c2 = "12346"
	c3 = "1"
	c4 = "0"
)

func TestBitDifferenceEquality(t *testing.T) {
	t.Parallel()

	a := sha256.Sum256([]byte(c1))
	b := sha256.Sum256([]byte(c1))

	want := 0
	got := bitdifference.BitDifference(a, b)

	if want != got {
		t.Errorf("wanted %d, but got %d", want, got)
	}
}

func TestBitDifferenceBy125(t *testing.T) {
	t.Parallel()

	a := sha256.Sum256([]byte(c3))
	b := sha256.Sum256([]byte(c4))

	want := 125
	got := bitdifference.BitDifference(a, b)

	if want != got {
		t.Errorf("wanted %d, but got %d", want, got)
	}
}
