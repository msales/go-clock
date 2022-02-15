package clock_test

import (
	"testing"
	"time"

	"github.com/msales/go-clock"
)

func TestRace_After(t *testing.T) {
	testMockInRace(func() {
		clock.After(1 * time.Second)
	})
}

func TestRace_AfterFunc(t *testing.T) {
	testMockInRace(func() {
		clock.AfterFunc(1*time.Second, func() {

		})
	})
}

func TestRace_Now(t *testing.T) {
	testMockInRace(func() {
		clock.Now()
	})
}

func TestRace_Since(t *testing.T) {
	testMockInRace(func() {
		clock.Since(time.Now())
	})
}

func TestRace_Sleep(t *testing.T) {
	testInRace(func() {
		clock.Sleep(1 * time.Nanosecond)
	})
}

func TestRace_Tick(t *testing.T) {
	testMockInRace(func() {
		clock.Tick(1 * time.Nanosecond)
	})
}

func TestRace_Timer(t *testing.T) {
	testMockInRace(func() {
		clock.Timer(1 * time.Nanosecond)
	})
}

func TestRace_Mock(t *testing.T) {
	testMockInRace(func() {
		clock.Mock(time.Now())
	})
}

func testMockInRace(runFunc func()) {
	now := time.Date(2019, time.September, 30, 14, 30, 00, 00, time.UTC)

	wait1 := make(chan struct{}, 1)
	wait2 := make(chan struct{}, 1)

	clock.Mock(now)
	go func() {
		runFunc()
		wait1 <- struct{}{}
		<-wait2
	}()

	go func() {
		runFunc()
		wait2 <- struct{}{}
		<-wait1
	}()

	clock.Restore()
}

func testInRace(runFunc func()) {
	wait1 := make(chan struct{}, 1)
	wait2 := make(chan struct{}, 1)

	go func() {
		runFunc()
		wait1 <- struct{}{}
		<-wait2
	}()

	go func() {
		runFunc()
		wait2 <- struct{}{}
		<-wait1
	}()
}
