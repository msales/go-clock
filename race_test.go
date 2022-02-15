package clock_test

import (
	"testing"
	"time"

	"github.com/msales/go-clock"
)

func TestRace(t *testing.T) {
	now := time.Date(2019, time.September, 30, 14, 30, 00, 00, time.UTC)

	wait1 := make(chan struct{}, 1)
	wait2 := make(chan struct{}, 1)

	go func() {
		clock.Mock(now)
		wait1 <- struct{}{}
		<-wait2
	}()

	go func() {
		clock.Mock(now)
		wait2 <- struct{}{}
		<-wait1
	}()

	clock.Restore()
}
