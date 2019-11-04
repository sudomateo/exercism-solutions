// Package reverse contains a solution to the Reverse String exercism.io exercise.
package reverse

// Reverse reverses a string.
func Reverse(s string) string {
	runes := []rune(s)
	var reverse []rune
	for i := len(runes) - 1; i >= 0; i-- {
		reverse = append(reverse, runes[i])
	}
	return string(reverse)
}
