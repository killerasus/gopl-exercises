// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 16.
//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "status: %s\nfetch: %v\n", resp.Status, err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stdout, "status: %s\n", resp.Status)
		b, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "status: %s\nfetch: reading %s: %v\n%d bytes read.", resp.Status, url, err, b)
			os.Exit(1)
		}
	}
}
