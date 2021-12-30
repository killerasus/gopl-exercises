package main

import (
	"crypto/sha256"
	"fmt"
	"killerasus/exerc4_1/bitdifference"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 3 {
		name := filepath.Base(os.Args[0])
		fmt.Println(name, "compares the bit differerece of the sha256 encryption of two entries")
		fmt.Println("Usage:\n\t", name, "value1", "value2")
		return
	}

	a := sha256.Sum256([]byte(os.Args[1]))
	b := sha256.Sum256([]byte(os.Args[2]))

	fmt.Printf("a = %x\nb = %x\nBit difference = %d\n", a, b, bitdifference.BitDifference(a, b))
}
