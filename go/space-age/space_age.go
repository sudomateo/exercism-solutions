// Package space contains a solution to the Space Age exercise.
package space

// Planet represents the name of a planet.
type Planet string

// The number of seconds planet Earth takes to perform one revolution around the sun.
const earthYear float64 = 31557600

// Age calculates your age depending on what planet you're from.
func Age(seconds float64, planet Planet) (age float64) {
	yearRatio := map[Planet]float64{
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Earth":   1,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}

	ratio, ok := yearRatio[planet]
	if !ok {
		return 0
	}

	return (seconds / (ratio * earthYear))
}
