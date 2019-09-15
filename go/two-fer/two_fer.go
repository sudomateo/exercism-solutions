// Package twofer contains a solution to the Two Fer exercise.
package twofer

import "fmt"

// ShareWith shares something with someone else
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
