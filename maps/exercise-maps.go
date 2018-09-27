package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// WordCount unexported
func WordCount(s string) map[string]int {
	wc := make(map[string]int)
	for _, word := range strings.Fields(s) {
		wc[word]++
	}
	return wc
}

func main() {
	wc.Test(WordCount)
}
