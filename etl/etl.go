package etl

import (
	"unicode"
)

//Transform calulcates new way of scrabble scores
func Transform(given map[int][]string) map[string]int {
	count := make(map[string]int)
	for val, result := range given {
		for _, char := range result {
			char := []rune(char)
			letter := unicode.ToLower(char[0])
			count[string(letter)] = val
		}
	}
	return count
}
