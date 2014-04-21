// Exercise: Maps
package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	
	for _, v := range strings.Fields(s) {
		m[v] += 1  // Start 0
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
