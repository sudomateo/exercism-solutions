// Package luhn contains a solution to the Luhn exercise
package luhn

// Valid determines if a given string is valid per the Luhn formula
func Valid(s string) bool {

	var sum, index int

	for i := len(s) - 1; i >= 0; i-- {
		chr := s[i]
		if chr == ' ' {
			continue
		}
		if chr < '0' || chr > '9' {
			return false
		}
		digit := int(chr - '0')
		if index%2 != 0 {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		index++
	}

	if index < 2 {
		return false
	}

	return sum%10 == 0
}
