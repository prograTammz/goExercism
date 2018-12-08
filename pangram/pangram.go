package pangram

import (
	"strings"
)

func IsPangram(s string) bool {
	ascii := 97
	data := strings.ToLower(s)
	for ascii <= 122 {
		if strings.ContainsRune(data, rune(ascii)) {
			ascii++
			continue

		} else {
			return false
		}

	}
	return true
}
