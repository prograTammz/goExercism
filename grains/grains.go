package grains

import (
	"errors"
	"math"
)

//unvalid represents an error since we can't start counting from zero which is
//represented as -1 or less than zero
var errUnvalid = errors.New("UnValid number either negative or false")

//Square The logic of the problem matches the logic of the binary number, but since
//we start count them from 0 the given number will be decremented by one.
func Square(n int) (uint64, error) {
	if n <= 0 || n > 64 {
		return 0, errUnvalid
	}
	return uint64(math.Pow(2, float64(n-1))), nil
}

//Total Counts the total of the number from 2^0 till 2^64
func Total() uint64 {
	var sum uint64
	for i := 0; i < 64; i++ {
		sum += uint64(math.Pow(2, float64(i)))
	}
	return sum
}
