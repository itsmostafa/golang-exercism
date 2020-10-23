package isogram

import (
	"unicode"
)

// IsIsogram checks if a word is an isogram.
func IsIsogram(phrase string) bool {
	set := make(map[rune]struct{})

	for _, character := range phrase {
		letter := unicode.ToLower(character)
		if unicode.IsLetter(letter) {
			if _, ok := set[letter]; ok {
				return false
			}
			set[letter] = struct{}{}
		}
	}
	return true
}
