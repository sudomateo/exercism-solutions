// Package isogram contains a solution to the Isogram exercise
package isogram

import (
	"unicode"
)

// IsIsogram determines if a string is an isogram
func IsIsogram(s string) bool {
	m := make(map[rune]bool)

	for _, r := range s {
		r = unicode.ToLower(r)
		if r == ' ' || r == '-' {
			continue
		}
		if m[r] {
			return false
		}
		m[r] = true
	}

	return true
}
