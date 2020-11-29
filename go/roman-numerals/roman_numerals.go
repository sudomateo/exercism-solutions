package romannumerals

import (
	"bytes"
	"fmt"
)

// ToRomanNumeral takes an integer and converts it to a Roman numeral.
func ToRomanNumeral(n int) (string, error) {
	if n <= 0 || n > 3000 {
		return "", fmt.Errorf("cannot convert %d", n)
	}

	thousands := n / 1000 % 10
	hundreds := n / 100 % 10
	tens := n / 10 % 10
	ones := n % 10

	var b bytes.Buffer

	for i := 0; i < thousands; i++ {
		b.WriteString("M")
	}

	hundred, err := convertDigit(hundreds, "C", "D", "M")
	if err != nil {
		return "", fmt.Errorf("could not convert digit %d", hundreds)
	}
	ten, err := convertDigit(tens, "X", "L", "C")
	if err != nil {
		return "", fmt.Errorf("could not convert digit %d", tens)
	}
	one, err := convertDigit(ones, "I", "V", "X")
	if err != nil {
		return "", fmt.Errorf("could not convert digit %d", ones)
	}

	b.WriteString(hundred)
	b.WriteString(ten)
	b.WriteString(one)

	return b.String(), nil
}

// convertDigit converts a single digit to its equivalent Roman numeral representation
// using specified strings for its ones, fives, and tens places.
func convertDigit(digit int, numeralOne, numeralFive, numeralTen string) (string, error) {
	if digit < 0 || digit > 9 {
		return "", fmt.Errorf("cannot convert %d", digit)
	}

	var b bytes.Buffer

	switch digit {
	case 1, 2, 3:
		for i := 0; i < digit; i++ {
			b.WriteString(numeralOne)
		}
	case 4:
		b.WriteString(numeralOne + numeralFive)
	case 5, 6, 7, 8:
		b.WriteString(numeralFive)
		for i := 0; i < digit-5; i++ {
			b.WriteString(numeralOne)
		}
	case 9:
		b.WriteString(numeralOne + numeralTen)
	}

	return b.String(), nil
}
