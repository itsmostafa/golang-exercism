package clock

import "fmt"

// Clock struct
type Clock struct {
	hour   int
	minute int
}

// New generates the given hour and minute to time
func New(hour, minute int) Clock {
	c := Clock{hour: hour, minute: minute}
	for c.minute < 0 {
		c.minute += 60
		c.hour--
	}
	for c.hour < 0 {
		c.hour += 24
	}

	c.hour += int(c.minute / 60)
	c.minute = c.minute % 60
	c.hour = c.hour % 24

	return c
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
