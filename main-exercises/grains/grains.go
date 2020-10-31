package grains

import "errors"

// Square calculates the number of grains
// on chessboard square.
func Square(position int) (uint64, error) {
	if position < 1 || position > 64 {
		return 0, errors.New("Only choose positions between 1 and 64")
	}
	return 1 << (position - 1), nil
}

// Total returns the total number of grains
// on a chessboard.
func Total() uint64 {
	return 1<<64 - 1
}
