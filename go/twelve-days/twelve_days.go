// Package twelve prints the lyrics to the song "The Twelve Days of Christmas".
package twelve

import (
	"bytes"
	"fmt"
)

// verse holds what gift was given on a specific day.
type verse struct {
	day  string
	gift string
}

var (
	intro = "On the %s day of Christmas my true love gave to me: "
	song  = []verse{
		{day: "first", gift: "a Partridge in a Pear Tree"},
		{day: "second", gift: "two Turtle Doves"},
		{day: "third", gift: "three French Hens"},
		{day: "fourth", gift: "four Calling Birds"},
		{day: "fifth", gift: "five Gold Rings"},
		{day: "sixth", gift: "six Geese-a-Laying"},
		{day: "seventh", gift: "seven Swans-a-Swimming"},
		{day: "eighth", gift: "eight Maids-a-Milking"},
		{day: "ninth", gift: "nine Ladies Dancing"},
		{day: "tenth", gift: "ten Lords-a-Leaping"},
		{day: "eleventh", gift: "eleven Pipers Piping"},
		{day: "twelfth", gift: "twelve Drummers Drumming"},
	}
)

// Song prints the lyrics to the entire song.
func Song() string {
	var b bytes.Buffer
	for n := 1; n <= 12; n++ {
		b.WriteString(Verse(n))
		if n != 12 {
			b.WriteString("\n")
		}
	}
	return b.String()
}

// Verse prints a single verse in the song.
func Verse(n int) string {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf(intro, song[n-1].day))
	for i := n - 1; i >= 0; i-- {
		if n > 1 && i == 0 {
			b.WriteString("and ")
		}
		b.WriteString(song[i].gift)
		if n > 1 && i != 0 {
			b.WriteString(", ")
		}
	}
	b.WriteString(".")
	return b.String()
}
