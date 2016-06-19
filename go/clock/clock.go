package clock

import "fmt"

const testVersion = 4

// Clock is an exported type that contains integer
type Clock int

// New creates and return a Clock object
func New(hour, minute int) Clock {
	c := Clock((hour * 60 + minute) % 1440)
	if c < 0 {
		c += 1440
	}
	return c
}

// Returns a string
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}

// Add minutes to a hour
func (c Clock) Add(minutes int) Clock {
	return New(0, int(c) + minutes)
}
