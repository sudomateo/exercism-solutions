// Package grains contains a solution to the Grains exercise.
package grains

import "errors"

var squareError = errors.New("square: invalid input")

// Square returns the number of grains for a given square
func Square(i int) (uint64, error) {
	if i < 1 || i > 64 {
		return 0, squareError
	}
	var grains uint64 = 1
	for i-1 != 0 {
		grains *= 2
		i--
	}
	return grains, nil
}

// Total returns the total number of grains for 64 squares
func Total() uint64 {
	var total, grains uint64 = 1, 1
	for i := 0; i < 65; i++ {
		grains *= 2
		total += grains
	}
	return total
}
