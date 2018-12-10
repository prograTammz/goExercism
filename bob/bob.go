// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) (string, int) {
	var letterCounter int
	var capitalCounter int
	var totalCounter int
	var lastGraphic rune
	for _, i := range remark {
		if unicode.IsLetter(rune(i)) {
			letterCounter++
			if unicode.IsUpper(rune(i)) {
				capitalCounter++
			}
		}
		if unicode.IsGraphic(rune(i)) && !unicode.IsSpace(rune(i)) {
			totalCounter++
			lastGraphic = rune(i)
		}
	}
	if totalCounter == 0 {
		return "Fine. Be that way!", 0
	}
	if letterCounter == 0 {
		if lastGraphic == '?' {
			return "Sure.", letterCounter
		}
		return "Whatever.", letterCounter
	}
	if lastGraphic == '?' {
		if capitalCounter < letterCounter/2 {
			return "Sure.", letterCounter
		}
		return "Calm down, I know what I'm doing!", letterCounter
	} else {
		if capitalCounter < letterCounter/2 {
			return "Whatever.", letterCounter
		}
		return "Whoa, chill out!", letterCounter
	}
}
