package prime

import "math"

func Nth(n int) (int, bool) {

	var nth int
	var primeCount int
	count := 1
	if n <= 0 {
		return 0, false
	}
	for primeCount != n {

		if IsPrime(count) {
			primeCount++
			nth = count
		}
		count++
	}
	return nth, true
}
func IsPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}
