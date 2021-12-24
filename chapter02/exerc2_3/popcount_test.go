package popcount_test

import (
	"popcount"
	"testing"
)

func BenchmarkPopCountUnroll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountUnroll(1234567890)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountLoop(1234567890)
	}
}
