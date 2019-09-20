// Package luhn contains a solution to the Luhn exercise
package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

// Valid determines if a given string is valid per the Luhn formula
func Valid(s string) bool {
	s = strings.Replace(s, " ", "", -1)

	if len(s) < 2 {
		return false
	}

	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}

	var sum int

	for i, j := len(s)-1, 0; i >= 0; i, j = i-1, j+1 {
		digit, err := strconv.Atoi(string(s[i]))
		if err != nil {
			return false
		}
		if j%2 != 0 {
			if digit*2 > 9 {
				digit = (digit * 2) - 9
			} else {
				digit = digit * 2
			}
		}
		sum += digit
	}

	return sum%10 == 0
}
