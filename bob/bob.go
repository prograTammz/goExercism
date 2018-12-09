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
	if totalCounter != 0 && letterCounter == 0 && lastGraphic != '?' {
		return "Whatever.", letterCounter
	}
	if totalCounter != 0 && letterCounter == 0 && lastGraphic == '?' {
		return "Sure.", letterCounter
	}
	if lastGraphic == '?' && capitalCounter < letterCounter/2 {
		return "Sure.", letterCounter
	}
	if lastGraphic == '?' && capitalCounter >= letterCounter/2 {
		return "Calm down, I know what I'm doing!", letterCounter
	}
	if lastGraphic != '?' && capitalCounter < letterCounter/2 {
		return "Whatever.", letterCounter
	}
	if lastGraphic != '?' && capitalCounter >= letterCounter/2 {
		return "Whoa, chill out!", letterCounter
	}
	return "", 0
}
