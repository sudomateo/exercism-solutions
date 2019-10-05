// Package letter contains a solution to the Parallel Letter Frequency exercise.
package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency calls the Frequency function concurrently.
func ConcurrentFrequency(xs []string) FreqMap {
	c := make(chan FreqMap)

	for _, v := range xs {
		go func(s string) {
			c <- Frequency(s)
		}(v)
	}

	result := FreqMap{}
	for range xs {
		for k, v := range <-c {
			result[k] += v
		}
	}

	return result
}
