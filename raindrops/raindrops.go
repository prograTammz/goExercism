package raindrops

import (
	"strconv"
	"strings"
)

func Convert(value int) string {
	result := strings.Builder{}
	if value%3 == 0 {
		result.WriteString("Pling")
	}

	if value%5 == 0 {
		result.WriteString("Plang")
	}

	if value%7 == 0 {
		result.WriteString("Plong")
	}
	finalResult := result.String()
	if len(finalResult) == 0 {
		finalResult = strconv.Itoa(value)
	}
	print(finalResult)
	return finalResult
}
