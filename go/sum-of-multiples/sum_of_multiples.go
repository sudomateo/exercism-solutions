package summultiples

// SumMultiples returns the sum of all numbers from 0 to limit that are
// divisible by at least one of the divisors.
func SumMultiples(limit int, divisors ...int) int {
	var sum int
	for i := 1; i < limit; i++ {
		for _, divisor := range divisors {
			if divisor == 0 {
				continue
			}
			if i%divisor == 0 {
				sum += i
				break
			}
		}
	}
	return sum
}
