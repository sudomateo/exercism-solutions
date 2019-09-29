// Package grains contains a solution to the Grains exercise.
package grains

import (
	"errors"
)

// Square returns the number of grains for a given square
func Square(i int) (uint64, error) {
	if i < 1 || i > 64 {
		return 0, errors.New("square: invalid input")
	}
	return uint64(1 << (i - 1)), nil
}

// Total returns the total number of grains for 64 squares
func Total() uint64 {
	return uint64((1 << 64) - 1)
}
