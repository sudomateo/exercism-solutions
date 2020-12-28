// Package anagram contains a solution to the anagram Exercism exercise.
package anagram

import (
	"sort"
	"strings"
)

// Detect determines whether or not any of the strings within cases are an
// anagram of subject, returning a slice of strings that are anagrams.
func Detect(subject string, cases []string) []string {
	anagrams := make([]string, 0, len(cases))

	// Sort the subject string (case-insensitive).
	subLower := strings.ToLower(subject)
	subSlice := strings.Split(subLower, "")
	sort.Slice(subSlice, func(i int, j int) bool { return subSlice[i] < subSlice[j] })
	subSorted := strings.Join(subSlice, "")

	for _, word := range cases {
		if len(subject) != len(word) {
			continue
		}

		wordLower := strings.ToLower(word)

		// Words cannot be anagrams of themselves.
		if strings.EqualFold(subLower, wordLower) {
			continue
		}

		// Sort the current string (case-insensitive).
		wordSlice := strings.Split(wordLower, "")
		sort.Slice(wordSlice, func(i int, j int) bool { return wordSlice[i] < wordSlice[j] })
		wordSorted := strings.Join(wordSlice, "")

		// The current word is not an anagram.
		if !strings.EqualFold(subSorted, wordSorted) {
			continue
		}

		anagrams = append(anagrams, word)
	}

	return anagrams
}
