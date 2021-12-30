package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var hashfunc string
var help bool

func init() {
	flag.StringVar(&hashfunc, "hash", "sha256", "the hashing function. Accepted values: sha256, sha384, sha512")
	flag.BoolVar(&help, "help", false, "prints this help")
}

func main() {

	flag.Parse()

	if help {
		name := filepath.Base(os.Args[0])
		fmt.Println(name, "returns the hash of the standard input given a hash function")
		flag.PrintDefaults()
		os.Exit(1)
	}

	hashfunc := strings.ToLower(hashfunc)

	switch hashfunc {
	case "sha256":
	case "sha384":
	case "sha512":
	default:
		fmt.Println(hashfunc, "is not a valid parameter. Check help for details.")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()
		entry := scanner.Text()

		if len(entry) == 0 {
			break
		}

		switch hashfunc {
		case "sha512":
			fmt.Printf("%x\n", sha512.Sum512([]byte(entry)))
		case "sha384":
			fmt.Printf("%x\n", sha512.Sum384([]byte(entry)))
		default:
			fmt.Printf("%x\n", sha256.Sum256([]byte(entry)))
		}
	}
}
