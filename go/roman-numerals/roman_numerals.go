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

	if thousands > 0 {
		for i := 0; i < thousands; i++ {
			b.WriteString("M")
		}
	}

	if hundreds > 0 {
		switch hundreds {
		case 1, 2, 3:
			for i := 0; i < hundreds; i++ {
				b.WriteString("C")
			}
		case 4:
			b.WriteString("CD")
		case 5, 6, 7, 8:
			b.WriteString("D")
			for i := 0; i < hundreds-5; i++ {
				b.WriteString("C")
			}
		case 9:
			b.WriteString("CM")
		}
	}

	if tens > 0 {
		switch tens {
		case 1, 2, 3:
			for i := 0; i < tens; i++ {
				b.WriteString("X")
			}
		case 4:
			b.WriteString("XL")
		case 5, 6, 7, 8:
			b.WriteString("L")
			for i := 0; i < tens-5; i++ {
				b.WriteString("X")
			}
		case 9:
			b.WriteString("XC")
		}
	}

	if ones > 0 {
		switch ones {
		case 1, 2, 3:
			for i := 0; i < ones; i++ {
				b.WriteString("I")
			}
		case 4:
			b.WriteString("IV")
		case 5, 6, 7, 8:
			b.WriteString("V")
			for i := 0; i < ones-5; i++ {
				b.WriteString("I")
			}
		case 9:
			b.WriteString("IX")
		}
	}

	return b.String(), nil
}
