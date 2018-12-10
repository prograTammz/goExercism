package collatzconjecture

import "errors"

var calError error = errors.New("un valid number")

func CollatzConjecture(n int) (int, error) {
	var counter int
	if n <= 0 {
		return 0, calError
	}
	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = n*3 + 1
		}
		counter++
	}
	return counter, nil
}
