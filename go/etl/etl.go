package etl

import "strings"

// Transform converts old scrabble scores into new scrabble scores.
func Transform(scores map[int][]string) map[string]int {
	newScores := make(map[string]int)
	for key, val := range scores {
		for _, letter := range val {
			lower := strings.ToLower(letter)
			newScores[lower] = key
		}
	}
	return newScores
}
