package goclock_test

import (
	"testing"

	goclock "github.com/msales/go-clock/v2"
)

func BenchmarkNopLogger_Info_Lock(b *testing.B) {
	goclock.UseLock()

	b.SetParallelism(100)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			goclock.Now()
		}
	})
}

func BenchmarkNopLogger_Info_NoLock(b *testing.B) {
	goclock.NoLock()

	b.SetParallelism(100)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			goclock.Now()
		}
	})
}
