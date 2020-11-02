package clock

import "fmt"

// Clock struct
type Clock struct {
	hour   int
	minute int
}

// New generates the given hour and minute to time
func New(hour, minute int) Clock {
	m := minute + hour*60
	m %= 24 * 60
	if m < 0 {
		m += 24 * 60
	}
	return Clock{m / 60, m % 60}
}

// Add the given duration.
func (c Clock) Add(duration int) Clock {
	c.minute += duration
	return New(c.hour, c.minute)
}

// Subtract the given duration.
func (c Clock) Subtract(duration int) Clock {
	c.minute -= duration
	return New(c.hour, c.minute)
}

// String converts the given time to a string.
func (c Clock) String() string {
	return fmt.Sprintf("%02v:%02v", c.hour, c.minute)
}
