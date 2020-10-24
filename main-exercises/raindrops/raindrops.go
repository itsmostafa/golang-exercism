package raindrops

import (
	"strconv"
)

// Convert will return raindrop sounds based on input number.
func Convert(number int) string {

	var res string
	if number%3 == 0 {
		res += "Pling"
	}
	if number%5 == 0 {
		res += "Plang"
	}
	if number%7 == 0 {
		res += "Plong"
	}
	if res == "" {
		return strconv.Itoa(number)
	}

	return res
}