package time

import "time"

// CurrentTime just return the current time nothing fancy
func CurrentTime(y string) string {
	x := time.Now().Format(time.UnixDate)
	return y + x
}
