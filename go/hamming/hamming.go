// Package hamming contains a solution to the Hamming exercise.
package hamming

import "fmt"

// Distance computes the hamming difference between DNA strands
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("left strand is size %d and right strand is size %d", len(a), len(b))
	}
	var count int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}
	return count, nil
}
