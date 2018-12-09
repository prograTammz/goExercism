package encode

import (
	"strconv"
	"unicode"
)

func RunLengthEncode(s string) string {
	if len(s) == 0 {
		return s
	}
	currentGraphic := rune(s[0])
	count := 1
	var output string
	for i := 1; i < len(s); i++ {
		if rune(s[i]) == currentGraphic {
			count++
		} else {
			if count == 1 {
				output += string(currentGraphic)
			} else {
				output += strconv.Itoa(count)
				output += string(currentGraphic)
			}
			currentGraphic = rune(s[i])
			count = 1
		}
	}
	if count == 1 {
		output += string(currentGraphic)
	} else {
		output += strconv.Itoa(count)
		output += string(currentGraphic)
	}
	return output
}
func RunLengthDecode(s string) string {
	var output string
	digit := "0"
	for _, i := range s {
		if unicode.IsDigit(rune(i)) {
			digit += string(rune(i))
		} else {
			counter, err := strconv.Atoi(digit)
			if err == nil {
				for j := 0; j < counter; j++ {
					output += string(rune(i))
				}
			}
			digit = "0"
		}
	}
	return output
}
