package main

import (
	"os"
	"strconv"
	"strings"
	"testing"
)

func BenchmarkStringComposition(b *testing.B) {
	for j := 0; j < b.N; j++ {
		var composed string
		for i, arg := range os.Args[0:] {
			composed += strconv.Itoa(i) + " " + arg
		}
	}
}

func BenchmarkStringJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Join(os.Args[0:], " ")
	}
}
