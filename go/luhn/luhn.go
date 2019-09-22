// Package luhn contains a solution to the Luhn exercise
package luhn

import (
	"strings"
)

// Valid determines if a given string is valid per the Luhn formula
func Valid(s string) bool {
	s = strings.Replace(s, " ", "", -1)

	if len(s) < 2 {
		return false
	}

	var sum int

	for i, j := len(s)-1, 0; i >= 0; i, j = i-1, j+1 {
		chr := s[i]
		if chr < 47 || chr > 58 {
			return false
		}

		digit := int(chr - '0')
		if j%2 != 0 {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}

	return sum%10 == 0
}
