package popcount_test

import (
	"popcount"
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(1234567890)
	}
}

func TestPopCount(t *testing.T) {
	want := popcount.PopCountUnroll(1234567890)
	got := popcount.PopCount(1234567890)
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}
