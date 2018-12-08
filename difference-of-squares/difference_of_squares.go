package diffsquares

//SquareOfSum calculates the sum of the values from zero to given number
//using guass's law
func SquareOfSum(value int) int {
	sum := value * (value + 1) / 2
	return sum * sum
}

//SumOfSquares uses a mathmatical induction to calculate the sum
//of the squares of the given value
func SumOfSquares(value int) int {
	sum := value * (value + 1) * (2*value + 1) / 6
	return sum
}

//Difference call SquareOfSum &SumOfSquares functions
// to get the difference between them
func Difference(value int) int {
	return SquareOfSum(value) - SumOfSquares(value)
}
