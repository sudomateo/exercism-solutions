// Package twofer contains my answer to life
// Package twofer contains a solution to the Two Fer exercise.
package twofer

import "fmt"

// ShareWith shares something with someone else
func ShareWith(name string) string {
	var s string
	if name != "" {
		s = fmt.Sprintf("One for %v, one for me.", name)
	} else {
		s = "One for you, one for me."
	}
	return s
}
