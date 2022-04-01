package goclock

import (
	"context"
	"sync"
	"time"

	"github.com/benbjohnson/clock"
)

const (
	// Day represents full day.
	Day = 24 * time.Hour
	// Week represents full week.
	Week = 7 * Day
)

var (
	// def is the package-level clock.Clock instance.
	def = newClk(clock.New())
)

func newClk(clk clock.Clock) *localClock {
	return &localClock{
		clock: clk,
	}
}

type localClock struct {
	// Used to sync overriding clock. Locking is disabled by Default
	mutex mutexWrap
	clock clock.Clock
}

func (l *localClock) After(d time.Duration) <-chan time.Time {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	return l.clock.After(d)
}

func (l *localClock) AfterFunc(d time.Duration, f func()) *clock.Timer {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	return l.clock.AfterFunc(d, f)
}

func (l *localClock) Now() time.Time {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	return l.clock.Now()
}

func (l *localClock) Since(t time.Time) time.Duration {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	return l.clock.Since(t)
}

func (l *localClock) Sleep(d time.Duration) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.clock.Sleep(d)
}

func (l *localClock) Tick(d time.Duration) <-chan time.Time {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	return l.clock.Tick(d)
}

func (l *localClock) Ticker(d time.Duration) *clock.Ticker {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	return l.clock.Ticker(d)
}

func (l *localClock) Timer(d time.Duration) *clock.Timer {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	return l.clock.Timer(d)
}

func (l *localClock) Until(t time.Time) time.Duration {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	return l.clock.Until(t)
}

func (l *localClock) WithDeadline(parent context.Context, d time.Time) (context.Context, context.CancelFunc) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	return l.clock.WithDeadline(parent, d)
}

func (l *localClock) WithTimeout(parent context.Context, t time.Duration) (context.Context, context.CancelFunc) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	return l.clock.WithTimeout(parent, t)
}

func (l *localClock) set(clk clock.Clock) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.clock = clk
}

func (l *localClock) get() clock.Clock {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	return l.clock
}

func (l *localClock) restore() {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.clock = clock.New()
}

func (l *localClock) setLock() {
	l.mutex.Enable()
}

func (l *localClock) setNoLock() {
	l.mutex.Disable()
}

// After waits for the duration to elapse and then sends the current time
func After(d time.Duration) <-chan time.Time {
	return def.After(d)
}

// AfterFunc waits for the duration to elapse and then calls f in its own goroutine.
func AfterFunc(d time.Duration, f func()) *clock.Timer {
	return def.AfterFunc(d, f)
}

// Now returns the current local time.
func Now() time.Time {
	return def.Now()
}

// Since returns the time elapsed since t.
func Since(t time.Time) time.Duration {
	return def.Since(t)
}

// Sleep pauses the current goroutine for at least the duration d.
func Sleep(d time.Duration) {
	def.Sleep(d)
}

// Tick is a convenience wrapper for NewTicker providing access to the ticking channel only.
func Tick(d time.Duration) <-chan time.Time {
	return def.Tick(d)
}

// Ticker returns a new Ticker containing a channel that will send the
// time with a period specified by the duration argument.
func Ticker(d time.Duration) *clock.Ticker {
	return def.Ticker(d)
}

// Timer creates a new Timer that will send the current time on its channel after at least duration d.
func Timer(d time.Duration) *clock.Timer {
	return def.Timer(d)
}

func Until(t time.Time) time.Duration {
	return def.Until(t)
}

func WithDeadline(parent context.Context, d time.Time) (context.Context, context.CancelFunc) {
	return def.WithDeadline(parent, d)
}

func WithTimeout(parent context.Context, t time.Duration) (context.Context, context.CancelFunc) {
	return def.WithTimeout(parent, t)
}

// Current returns current instance of clock
func Current() clock.Clock {
	return def.get()
}

// Mock sets the Clock with a mock frozen at the given time and returns it.
func Mock(now time.Time) *clock.Mock {
	mock := clock.NewMock()
	mock.Set(now)

	def.set(mock)

	return mock
}

// Set sets the clock.
func Set(clk clock.Clock) {
	def.set(clk)
}

// Restore replaces the Clock with the real clock.
func Restore() {
	def.restore()
}

// UseLock adds locking mechanism back on clock.
func UseLock() {
	def.setLock()
}

// NoLock removes locking mechanism usage on clock.
func NoLock() {
	def.setNoLock()
}

type mutexWrap struct {
	lock     sync.Mutex
	disabled bool
}

func (mw *mutexWrap) Lock() {
	if !mw.disabled {
		mw.lock.Lock()
	}
}

func (mw *mutexWrap) Unlock() {
	if !mw.disabled {
		mw.lock.Unlock()
	}
}

func (mw *mutexWrap) Enable() {
	mw.lock.Lock()
	defer mw.lock.Unlock()

	mw.disabled = false
}

func (mw *mutexWrap) Disable() {
	mw.lock.Lock()
	defer mw.lock.Unlock()

	mw.disabled = true
}
