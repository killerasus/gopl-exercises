package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

// comma inserts commas in a floating point string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	start := 0
	if s[0] == '+' || s[0] == '-' {
		start = 1
	}

	floating := strings.LastIndex(s, ".")

	var commatized bytes.Buffer
	for i, j := n-1, 0; i >= start; i-- {
		if j >= 3 && i < floating {
			commatized.WriteByte(',')
			j = 0
		}
		commatized.WriteByte(s[i])
		if i < floating {
			j++
		}
	}

	reversed := commatized.Bytes()
	var ret bytes.Buffer
	if start > 0 {
		ret.WriteByte(s[0])
	}
	for i := len(reversed) - 1; i >= 0; i-- {
		ret.WriteByte(reversed[i])
	}

	return ret.String()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:\n\tcomma <number>")
		os.Exit(1)
	}
	fmt.Println(comma(os.Args[1]))
}
