package scrabble

import "unicode"

// Score calculates the scrabble value for the input word.
func Score(word string) int {
	var value int
	for _, letter := range word {
		switch unicode.ToUpper(letter) {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			value++
		case 'D', 'G':
			value += 2
		case 'B', 'C', 'M', 'P':
			value += 3
		case 'F', 'H', 'V', 'W', 'Y':
			value += 4
		case 'K':
			value += 5
		case 'J', 'X':
			value += 8
		case 'Q', 'Z':
			value += 10
		}
	}
	return value
}
