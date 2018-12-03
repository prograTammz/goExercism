package scrabble

import (
	"unicode"
)

var letter = map[rune]int{
	'a': 1,
	'e': 1,
	'i': 1,
	'o': 1,
	'u': 1,
	'l': 1,
	'n': 1,
	'r': 1,
	's': 1,
	't': 1,
	//the value "2" letters
	'd': 2,
	'g': 2,
	//the value 3 letters
	'b': 3,
	'c': 3,
	'm': 3,
	'p': 3,
	//the value 4 letters
	'f': 4,
	'h': 4,
	'v': 4,
	'w': 4,
	'y': 4,
	//the value 5 letters
	'k': 5,
	//the value 8 letters
	'j': 8,
	'x': 8,
	//the value 10 letter
	'q': 10,
	'z': 10,
}

//Score functions loops over each word and calculate it's score from the
//contained letters at which each letter has a score.
func Score(word string) int {

	var result []rune
	for _, val := range word {
		result = append(result, unicode.ToLower(val))
	}
	var count int
	for _, char := range result {
		count += letter[char]
	}
	return count
}
