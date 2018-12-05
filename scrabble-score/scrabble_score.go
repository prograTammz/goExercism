package scrabble

import (
	"unicode"
)

//Score functions loops over each word and calculate it's score from the
//contained letters at which each letter has a score.
func Score(word string) int {

	var count int
	for _, val := range word {
		result := unicode.ToLower(val)
		switch result {
		case 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
			count++
		case 'd', 'g':
			count += 2
		case 'b', 'c', 'm', 'p':
			count += 3
		case 'f', 'h', 'v', 'w', 'y':
			count += 4
		case 'k':
			count += 5
		case 'j', 'x':
			count += 8
		case 'q', 'z':
			count += 10
		}

	}
	return count
}
