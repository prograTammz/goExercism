package secret

import (
	"strconv"
)

func Handshake(n uint) []string {
	var s string
	var reversedString string
	var final []string
	var isReverse bool
	for n >= 1 {
		if n%2 == 0 {
			s += "0"
		} else {
			s += "1"

		}
		n /= 2
	}
	reversedString = reverseString(s)
	val, err := strconv.Atoi(reversedString)
	if val >= 10000 {
		val -= 10000
		isReverse = true
	}
	if val >= 1000 {
		val -= 1000
		final = append(final, "jump")
	}
	if val >= 100 {
		val -= 100
		final = append(final, "close your eyes")
	}
	if val >= 10 {
		val -= 10
		final = append(final, "double blink")
	}
	if val >= 1 {
		val--
		final = append(final, "wink")
	}
	if err != nil {
		return nil
	}
	if !isReverse {
		return reverseStringArr(final)
	}
	return final
}
func reverseStringArr(sarr []string) []string {
	var result []string
	for i := len(sarr) - 1; i >= 0; i-- {
		result = append(result, sarr[i])
	}
	return result
}
func reverseString(s string) string {
	var rs string
	for i := len(s) - 1; i >= 0; i-- {
		rs += string(s[i])
	}
	return rs
}
