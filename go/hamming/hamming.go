package hamming

import "fmt"

func Distance(a, b string) (int, error) {
	var count int
	if len(a) == len(b) {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				count++
			}
		}
	} else {
		return 0, fmt.Errorf("Error: left strand is size %d, right strand is size %d", len(a), len(b))
	}
	return count, nil
}
