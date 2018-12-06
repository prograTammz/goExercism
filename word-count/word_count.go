// Package wordcount implements the wordcount exercise.
package wordcount

import (
	"strings"
	"unicode"
)

// Frequency relates words to their count.
type Frequency map[string]int

// WordCount counts the frequency of unique words in a string.
func WordCount(phrase string) Frequency {
	freq := make(Frequency)
	phrase = strings.ToLower(phrase)
	// split on nonapostrophe punctuation, spaces, and symbols
	words := strings.FieldsFunc(phrase, func(r rune) bool {
		return r != '\'' && (unicode.IsPunct(r) || unicode.IsSpace(r) || unicode.IsSymbol(r))
	})
	for _, w := range words {
		// but if words is surrounded by apostrophes then ignore them
		if len(w) > 2 && w[0] == '\'' && w[len(w)-1] == '\'' {
			freq[w[1:len(w)-1]]++
		} else {
			freq[w]++
		}
	}
	return freq
}
b