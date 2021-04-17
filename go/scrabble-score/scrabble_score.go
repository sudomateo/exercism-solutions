// Package scrabble helps win the game of scrabble.
package scrabble

import (
	"strings"
)

var pointMap = map[string]int{
	"AEIOULNRST": 1,
	"DG":         2,
	"BCMP":       3,
	"FHVWY":      4,
	"K":          5,
	"JX":         8,
	"QZ":         10,
}

// Score caclulates a scrabble score for a given string.
func Score(s string) int {
	var score int
	s = strings.ToUpper(s)
	for i := 0; i < len(s); i++ {
		for letters, points := range pointMap {
			if strings.Contains(letters, string(s[i])) {
				score += points
			}
		}
	}
	return score
}
