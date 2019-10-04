// Package clock contains a solution to the Clock exercism.io exercise.
package clock

import (
	"fmt"
)

const (
	minutesPerDay int = 1440
)

// Clock is the time in minutes
type Clock struct {
	Minutes int
}

// New contructs a new Clock.
func New(h, m int) Clock {
	totalMinutes := ((h * 60) + m) % minutesPerDay
	if totalMinutes < 0 {
		totalMinutes = minutesPerDay + totalMinutes
	}
	return Clock{
		Minutes: totalMinutes,
	}
}

// String prints Clock in a specific format
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Minutes/60, c.Minutes%60)

}

// Add adds time to a Clock and returns a new Clock
func (c Clock) Add(m int) Clock {
	return New(0, c.Minutes+m)
}

// Subtract subtracts time from a Clock and returns a new Clock
func (c Clock) Subtract(m int) Clock {
	return New(0, c.Minutes-m)
}
