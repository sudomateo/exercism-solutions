//Package raindrops is like FizzBuzz, but not.
package raindrops

import "strconv"

//Convert outputs strings based on an integer's factors.
func Convert(i int) string {
	var s string
	if i%3 == 0 {
		s += "Pling"
	}
	if i%5 == 0 {
		s += "Plang"
	}
	if i%7 == 0 {
		s += "Plong"
	}
	if s == "" {
		s = strconv.Itoa(i)
	}
	return s
}
