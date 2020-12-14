package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of a list of runes in parallel
func ConcurrentFrequency(text []string) FreqMap {
	c := make(chan FreqMap, len(text))
	m := FreqMap{}
	for _, s := range text {
		str := s
		go func() {
			c <- Frequency(str)
		}()
		for char, freq := range <-c {
			m[char] += freq
		}
	}

	return m
}
