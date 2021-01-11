package clock

import "time"

var now = time.Now().UTC

func Now() time.Time {
	return now()
}

func SetFakeTimeNow(fakeTimeNow func() time.Time) {
	now = fakeTimeNow
}
