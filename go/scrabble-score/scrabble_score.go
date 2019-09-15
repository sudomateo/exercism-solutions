// Package scrabble helps win the game of scrabble
package scrabble

import (
	"strings"
)

var points = map[string]int{
	"AEIOULNRST": 1,
	"DG":         2,
	"BCMP":       3,
	"FHVWY":      4,
	"K":          5,
	"JX":         8,
	"QZ":         10,
}

// Score caclulates my score for a given string
func Score(s string) int {
	if s == "" {
		return 0
	}
	var score int
	s = strings.ToUpper(s)
	for _, c := range s {
		for i, v := range points {
			if strings.Contains(i, string(c)) {
				score += v
			}
		}
	}
	return score
}
