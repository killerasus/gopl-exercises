package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int) // counts of Unicode characters
	letters := make(map[rune]int)
	digits := make(map[rune]int)
	controls := make(map[rune]int)
	punctuations := make(map[rune]int)
	spaces := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		counts[r]++
		utflen[n]++

		if unicode.IsLetter(r) {
			letters[r]++
		}
		if unicode.IsDigit(r) {
			digits[r]++
		}
		if unicode.IsControl(r) {
			controls[r]++
		}
		if unicode.IsPunct(r) {
			punctuations[r]++
		}
		if unicode.IsSpace(r) {
			spaces[r]++
		}
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("\nletters\tcount\n")
	for l, n := range letters {
		fmt.Printf("%c\t%d\n", l, n)
	}
	fmt.Printf("\ndigits\tcount\n")
	for d, n := range digits {
		fmt.Printf("%c\t%d\n", d, n)
	}
	fmt.Printf("\ncontrols\tcount\n")
	for c, n := range controls {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("\npuncts\tcount\n")
	for p, n := range punctuations {
		fmt.Printf("%c\t%d\n", p, n)
	}
	fmt.Printf("\nspaces\tcount\n")
	for s, n := range controls {
		fmt.Printf("%q\t%d\n", s, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
