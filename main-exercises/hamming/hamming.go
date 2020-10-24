package hamming

import "errors"

// Distance calculates the hamming distance between two inputs (DNA strands).
func Distance(a, b string) (int, error) {
	ar, br := []rune(a), []rune(b)

	if len(ar) != len(br) {
		return 0, errors.New("inputs are not the same length")
	}
	var distance int
	for i := range ar {
		if ar[i] != br[i] {
			distance++
		}
	}

	return distance, nil
}
