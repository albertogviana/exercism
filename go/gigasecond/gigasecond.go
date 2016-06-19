package gigasecond

import "time"

const testVersion = 4 // find the value in gigasecond_test.go

// AddGigasecond calculate when will be your gigasecond
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1e9)
}
