package diffsquares

// SquareOfSum calculates the squared sum of the input's
// first natural numbers.
func SquareOfSum(number int) int {
	sum := number * (number + 1) / 2
	return sum * sum
}

// SumOfSquares calculates the sum of the input's
// first natural numbers squared.
func SumOfSquares(number int) int {
	return number * (number + 1) * (2*number + 1) / 6
}

// Difference calculates the difference between
// Square of Sum and Sum of Squares.
func Difference(number int) int {
	return SquareOfSum(number) - SumOfSquares(number)
}
