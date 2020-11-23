package livebuilder

import "time"

// scheduleDelayed takes a function and waits at least the given delay before invoking the given function.
func scheduleDelayed(delay time.Duration, f func()) chan bool {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-time.After(delay):
			case <-stop:
				return
			}

			f()
		}
	}()

	return stop
}

