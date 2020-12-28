// Package wordcount contains a solution to the Exercism word-count exercise.
package wordcount

import (
	"regexp"
	"strings"
)

// Frequency is the frequency that words appears in a phrase.
type Frequency map[string]int

// WordCount counts the number of words within a given phrase, returning a
// Frequency value.
func WordCount(phrase string) Frequency {
	re := regexp.MustCompile(`[^a-zA-Z0-9']`)
	frequency := make(Frequency)

	p := re.ReplaceAllString(strings.ToLower(phrase), " ")

	for _, word := range strings.Fields(p) {
		word = strings.Trim(word, `'`)
		frequency[word]++
	}

	return frequency
}
