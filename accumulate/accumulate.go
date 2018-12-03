package accumulate

func Accumulate(given []string, converter func(string) string) []string {
	var result []string
	for _, char := range given {
		result = append(result, converter(char))
	}
	return result
}
