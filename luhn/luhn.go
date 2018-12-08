package luhn

import (
	"unicode"
)

//Valid Function takes a CC numebr or Candian ssn and returns if valid or no !
func Valid(s string) bool {
	count := 1
	digitCount := 0
	var sum int
	if len(s) == 0 {
		return false
	}
	for i := len(s) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(s[i])) {
			digitCount++
			tempVal := (int(s[i]) - '0')
			if count%2 == 0 {
				tempVal *= 2
				if tempVal > 9 {
					tempVal -= 9
				}
			}
			sum += tempVal
			count++
		}
		if !unicode.IsDigit(rune(s[i])) && !unicode.IsSpace(rune(s[i])) {
			return false
		}
	}
	if digitCount <= 1 {
		return false
	}
	if sum%10 == 0 {
		return true
	}
	return false
}
