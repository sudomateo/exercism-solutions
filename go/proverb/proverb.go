// Package proverb contains a solution to the Proverb exercise
package proverb

import (
	"fmt"
)

// Proverb tells a story given a slice of strings
func Proverb(rhyme []string) []string {
	var proverb []string

	for i, _ := range rhyme {
		if i < (len(rhyme) - 1) {
			proverb = append(proverb, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
			continue
		}
		proverb = append(proverb, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	}

	return proverb
}
