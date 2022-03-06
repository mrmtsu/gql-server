package adapter

import "time"

func NewSystemNowFunc() func() time.Time {
	return time.Now
}
