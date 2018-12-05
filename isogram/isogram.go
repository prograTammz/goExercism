package isogram

import (
	"unicode"
)

//IsIsogram function checks if there is repeatitive letter in the given word or not
func IsIsogram(input string) bool {

	seen := map[rune]bool{}
	for _, i := range input {
		i := unicode.ToLower(i)
		if !unicode.IsLetter(i) {
			continue
		}
		isThere := seen[i]
		if isThere {
			return false
		}
		seen[i] = true
	}
	return true
}
