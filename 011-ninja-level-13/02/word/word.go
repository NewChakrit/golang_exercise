// Package word provides custom function for working with strings and words
package word

import (
	"strings"
)

// UseCount returns the number of times each word is used in a string
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count returns th number of word in a string
func Count(s string) int {
	sp := strings.Fields(s)
	return len(sp)
}
