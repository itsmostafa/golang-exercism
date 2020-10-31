package clock

import "fmt"

type Clock struct {
	hour   int
	minute int
}

// Update converts hour and minute inputs to time
func (c Clock) Update() Clock {
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

// New generates a new time.
func New(hour, minute int) Clock {
	c := Clock{hour: hour, minute: minute}
	return c.Update()
}

// Add the given duration.
func (c Clock) Add(duration int) Clock {
	c.minute += duration
	return c.Update()
}

// Subtract the given duration.
func (c Clock) Subtract(duration int) Clock {
	c.minute -= duration
	return c.Update()
}

// String converts the given time to a string.
func (c Clock) String() string {
	return fmt.Sprintf("%02v:%02v", c.hour, c.minute)
}
