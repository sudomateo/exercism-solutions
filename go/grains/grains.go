// Package grains contains a solution to the Grains exercise.
package grains

import (
	"errors"
	"log"
)

// Square returns the number of grains for a given square
func Square(i int) (uint64, error) {
	if i < 1 || i > 64 {
		return 0, errors.New("square: invalid input")
	}
	grains := 1
	for i-1 != 0 {
		grains *= 2
		i--
	}
	return uint64(grains), nil
}

// Total returns the total number of grains for 64 squares
func Total() uint64 {
	var total uint64
	for i := 1; i < 65; i++ {
		grains, err := Square(i)
		if err != nil {
			log.Println("total: error squaring number")
		}
		total += grains
	}
	return total
}
