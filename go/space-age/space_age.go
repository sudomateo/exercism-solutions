// Package space contains a solution to the Space Age exercise.
package space

type Planet string

const earthYearSeconds float64 = 31557600

// Age calculates your age depending on what planet you're from.
func Age(seconds float64, planet Planet) (age float64) {
	var yearRatio float64

	switch planet {
	case "Earth":
		yearRatio = 1
	case "Mercury":
		yearRatio = 0.2408467
	case "Venus":
		yearRatio = 0.61519726
	case "Mars":
		yearRatio = 1.8808158
	case "Jupiter":
		yearRatio = 11.862615
	case "Saturn":
		yearRatio = 29.447498
	case "Uranus":
		yearRatio = 84.016846
	case "Neptune":
		yearRatio = 164.79132
	}

	age = (seconds / (yearRatio * earthYearSeconds))

	return age
}
