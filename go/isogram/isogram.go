// Package isogram contains a solution to the Isogram exercise
package isogram

import (
	"strings"
)

// IsIsogram determines if a string is an isogram
func IsIsogram(s string) bool {
	s = strings.ToLower(s)
	m := make(map[rune]int)

	for _, r := range s {
		if string(r) == " " || string(r) == "-" {
			continue
		}
		m[r]++
		if m[r] > 1 {
			return false
		}
	}

	return true
}
