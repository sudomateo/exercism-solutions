// Package diffsquares contains a solution to the Difference of Squares exercise.
package diffsquares

// SquareOfSum calculates the square of a sum of integers
func SquareOfSum(n int) int {
	return ((n * (n + 1)) / 2) * ((n * (n + 1)) / 2)
}

// SumOfSquares calculates the sum of the squares of integers
func SumOfSquares(n int) int {
	return ((n * (n + 1) * ((2 * n) + 1)) / 6)
}

// Difference calculates the difference between the output of SquareOfSum and
// the output of SumOfSquares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
