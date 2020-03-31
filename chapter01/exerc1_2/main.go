package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for i, arg := range os.Args[0:] {
		fmt.Println(strconv.Itoa(i) + " " + arg)
	}
}
