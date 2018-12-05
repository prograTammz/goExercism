package reverse

import (
	"unicode/utf8"
)

//String flips the word
func String(input string) string {
	tempString := input
	var output string
	for len(tempString) > 0 {
		r, size := utf8.DecodeLastRuneInString(tempString)
		output += string(r)
		tempString = tempString[:len(tempString)-size]

	}
	return output
}
