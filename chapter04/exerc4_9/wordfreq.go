package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("wordfreq computes word frequency in a text file.")
		fmt.Println("Usage:\n\t'wordfreq <filename>'")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	wordfreq := make(map[string]int)

	for scanner.Scan() {
		wordfreq[strings.Trim(scanner.Text(), "_:-,!.?;\"''()")]++
	}

	fmt.Println("\nWord\tFrequency")
	for k, v := range wordfreq {
		fmt.Printf("%s\t%d\n", k, v)
	}
}
