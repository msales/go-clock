package goclock_test

import (
	"testing"

	goclock "github.com/msales/go-clock/v2"
)

func BenchmarkClock_Now_Lock(b *testing.B) {
	b.SetParallelism(100)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			goclock.Now()
		}
	})
}

func BenchmarkClock_Now_NoLock(b *testing.B) {
	goclock.NoLock()

	b.SetParallelism(100)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			goclock.Now()
		}
	})
}
