// Package clock contains a solution to the Clock exercism.io exercise.
package clock

import (
	"fmt"
)

// Clock is the time represented by two integers
type Clock struct {
	Hour   int
	Minute int
}

// New contructs a new Clock. This needs to be more efficient.
func New(h, m int) Clock {

	h = ((m / 60) + h) % 24
	m = m % 60

	if h < 0 && m >= 0 {
		h = 24 + h
	}

	if h < 0 && m < 0 {
		h = 23 + h
		m = 60 + m
	}

	if h >= 0 && m < 0 {
		h = (0 + h) - 1
		if h < 0 {
			h = 23
		}
		m = 60 + m
	}

	return Clock{
		Hour:   h,
		Minute: m,
	}
}

// String prints Clock in a specific format
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hour, c.Minute)
}

// Add adds time to a Clock and returns a new Clock
func (c Clock) Add(m int) Clock {
	return New(c.Hour, c.Minute+m)
}

// Substract subtracts time from a Clock and returns a new Clock
func (c Clock) Subtract(m int) Clock {
	return New(c.Hour, c.Minute-m)
}
