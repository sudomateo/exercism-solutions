// Package diffsquares contains a solution to the Difference of Squares exercise.
package diffsquares

// SquareOfSum calculates the square of a sum of integers
func SquareOfSum(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

// SumOfSquares calculates the sum of the squares of integers
func SumOfSquares(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}

// Difference calculates the difference between the output of SquareOfSum and
// the output of SumOfSquares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
