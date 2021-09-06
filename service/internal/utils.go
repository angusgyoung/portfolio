package internal

import "time"

func toSecondDuration(seconds int) time.Duration {
	return time.Duration(seconds * int(time.Second))
}