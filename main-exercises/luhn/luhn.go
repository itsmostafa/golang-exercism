package luhn

import (
	"strconv"
	"strings"
)

// Valid returns whether a string is passes the luhn formula.
func Valid(cc string) bool {
	cc = strings.ReplaceAll(cc, " ", "")

	if len(cc) <= 1 {
		return false
	}

	double := len(cc)%2 != 0

	var sum int
	for _, r := range cc {
		digit, err := strconv.Atoi(string(r))
		if err != nil {
			return false
		}
		double = !double

		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
	}

	return sum%10 == 0
}
