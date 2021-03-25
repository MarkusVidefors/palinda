package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	split := strings.Fields(s)
	count := make(map[string]int)

	for _, word := range split {
		count[word] += 1
	}

	return count
}

func main() {
	wc.Test(WordCount)
}

