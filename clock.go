package clock

import (
	"time"

	"github.com/benbjohnson/clock"
)

// init initializes the Clock variable with a real Clock.
func init() {
	Restore()
}

// Day represents full day.
const Day = 24 * time.Hour

// Clock represents a global clock.
var Clock clock.Clock

// After waits for the duration to elapse and then sends the current time
func After(d time.Duration) <-chan time.Time {
	return Clock.After(d)
}

// AfterFunc waits for the duration to elapse and then calls f in its own goroutine.
func AfterFunc(d time.Duration, f func()) *clock.Timer {
	return Clock.AfterFunc(d, f)
}

// Now returns the current local time.
func Now() time.Time {
	return Clock.Now()
}

// Since returns the time elapsed since t.
func Since(t time.Time) time.Duration {
	return Clock.Since(t)
}

// Sleep pauses the current goroutine for at least the duration d.
func Sleep(d time.Duration) {
	Clock.Sleep(d)
}

// Tick is a convenience wrapper for NewTicker providing access to the ticking channel only.
func Tick(d time.Duration) <-chan time.Time {
	return Clock.Tick(d)
}

// Ticker returns a new Ticker containing a channel that will send the
// time with a period specified by the duration argument.
func Ticker(d time.Duration) *clock.Ticker {
	return Clock.Ticker(d)
}

// Timer creates a new Timer that will send the current time on its channel after at least duration d.
func Timer(d time.Duration) *clock.Timer {
	return Clock.Timer(d)
}

// Mock replaces the Clock with a mock frozen at the given time and returns it.
func Mock(now time.Time) *clock.Mock {
	mock := clock.NewMock()
	mock.Set(now)

	Clock = mock

	return mock
}

// Restore replaces the Clock with the real clock.
func Restore() {
	Clock = clock.New()
}
