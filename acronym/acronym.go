// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	tempString := strings.FieldsFunc(s, Split)
	var output string
	for _, val := range tempString {
		output += string(unicode.ToUpper(rune(val[0])))
	}
	return output
}
func Split(x rune) bool {
	return x == '-' || x == ' '
}
