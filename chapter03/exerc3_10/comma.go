package main

import (
	"bytes"
	"fmt"
	"os"
)

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	var commatized bytes.Buffer
	for i, j := n-1, 0; i >= 0; i, j = i-1, j+1 {
		if j >= 3 {
			commatized.WriteByte(',')
			j = 0
		}
		commatized.WriteByte(s[i])
	}

	reversed := commatized.Bytes()
	var ret bytes.Buffer
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
