// Package triangle contains a solution to the Triangle exercise.
package triangle

import "math"

type Kind int

const (
	NaT = iota
	Equ = iota
	Iso = iota
	Sca = iota
)

// KindFromSides determines the type of triangle from its sides
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	switch {
	case (math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c)) ||
		(math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0)) ||
		(a <= 0 || b <= 0 || c <= 0) ||
		(a+b < c || b+c < a || c+a < b):
		k = NaT
	case a == b && b == c:
		k = Equ
	case a == b || b == c || c == a:
		k = Iso
	case a != b && b != c && c != a:
		k = Sca
	}
	return k
}
