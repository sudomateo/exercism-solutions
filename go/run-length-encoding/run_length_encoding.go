package encode

import (
	"strconv"
	"strings"
)

// writeEncoded writes a rune r to strings.Builder b count times.
func writeEncoded(b *strings.Builder, count int, r rune) {
	switch {
	case count == 1:
		b.WriteRune(r)
	case count > 1:
		b.WriteString(strconv.Itoa(count))
		b.WriteRune(r)
	}
}

// RunLengthEncode encodes a string via run length encoding.
func RunLengthEncode(input string) string {
	var count int
	var previous rune
	var b strings.Builder

	for _, r := range input {
		switch {
		case r == previous:
			count++
		default:
			writeEncoded(&b, count, previous)
			count = 1
			previous = r
		}
	}
	writeEncoded(&b, count, previous)
	return b.String()
}

// RunLengthDecode decodes a string that was encoded via run length encoding.
func RunLengthDecode(input string) string {
	var count int
	var b strings.Builder

	for _, r := range input {
		switch {
		case '0' <= r && r <= '9':
			count = count*10 + int(r-'0')
		default:
			if count == 0 {
				count = 1
			}
			for ; count > 0; count-- {
				b.WriteRune(r)
			}
		}
	}
	return b.String()
}
