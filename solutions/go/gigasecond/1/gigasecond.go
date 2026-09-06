package gigasecond

import "time"

func AddGigasecond(t time.Time) time.Time {
	gigasecond := time.Second * 1e9
	return t.Add(gigasecond)
}
