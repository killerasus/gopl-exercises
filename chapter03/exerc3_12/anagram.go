package main

import (
	"fmt"
	"os"
)

func IsAnagram(s1, s2 string) bool {
	s1map := make(map[rune]int)
	for _, c := range s1 {
		s1map[c]++
	}

	s2map := make(map[rune]int)
	for _, c := range s2 {
		s2map[c]++
	}

	for k := range s1map {
		if s1map[k] != s2map[k] {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println(os.Args[0], "verifies if two strings are anagrams")
		fmt.Println("Usage:\n\t", os.Args[0], "STRING1 STRING2")
		os.Exit(1)
	}

	fmt.Println("Are", os.Args[1], "and", os.Args[2], "anagrams?")
	fmt.Println(IsAnagram(os.Args[1], os.Args[2]))
}
