// Package clock contains a solution to the Clock exercism.io exercise.
package clock

import (
	"fmt"
)

const (
	hourInMinutes int = 60
	dayInMinutes  int = 24 * hourInMinutes
)

// Clock is the time in minutes.
type Clock struct {
	minutes int
}

// New contructs a new Clock.
func New(h, m int) Clock {
	return Clock{
		minutes: parseMinutes((h * hourInMinutes) + m),
	}
}

// String prints Clock in a specific format.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes/hourInMinutes, c.minutes%hourInMinutes)
}

// Add adds time to a Clock and returns a new Clock.
func (c Clock) Add(m int) Clock {
	return New(0, c.minutes+m)
}

// Subtract subtracts time from a Clock and returns a new Clock.
func (c Clock) Subtract(m int) Clock {
	return New(0, c.minutes-m)
}

// parseMinutes ensures minutes are non-negative and fall within one day.
func parseMinutes(m int) int {
	m %= dayInMinutes
	if m < 0 {
		m += dayInMinutes
	}
	return m
}
