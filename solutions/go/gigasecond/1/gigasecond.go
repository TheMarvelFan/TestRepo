// Package gigasecond contains a method add a gigasecond to a given date.
package gigasecond

import "time"

// AddGigasecond returns the time after adding a gigasecond to it.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1_000_000_000 * time.Second)
}
