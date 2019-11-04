// Package pangram contains a solution to the Pangram exercism.io exercise.
package pangram

import (
	"strings"
)

// IsPangram tests whether or not a given string is a pangram.
func IsPangram(s string) bool {
	seen := make(map[rune]bool)
	for i := 0; i < 26; i++ {
		seen[rune('a'+i)] = false
	}
	s = strings.ToLower(s)
	for _, r := range s {
		if r < 'a' || r > 'z' {
			continue
		}
		seen[r] = true
	}
	for _, b := range seen {
		if !b {
			return false
		}
	}
	return true
}
