// Package accumulate contains a solution to the Accumulate exercise.
package accumulate

// Accumulate performs a given operation on every element of a slice of strings.
func Accumulate(xs []string, f func(string) string) []string {
	for i, v := range xs {
		xs[i] = f(v)
	}
	return xs
}
