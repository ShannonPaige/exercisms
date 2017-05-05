package gigasecond

import "time"

const testVersion = 4

func AddGigasecond(t time.Time) time.Time {
	return t.Add(1E9 * time.Second)
}
