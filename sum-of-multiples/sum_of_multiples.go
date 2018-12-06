package summultiples

func SumMultiples(n int, multiples ...int) int {
	var sum int
	i := 0
	for i < n {
		for _, val := range multiples {
			if i%val == 0 {
				sum += i
				break
			}
		}
		i++
	}
	return sum
}
