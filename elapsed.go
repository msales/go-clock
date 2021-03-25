package clock

import (
	"math"
	"time"
)

// DayElapsed elapsed time during a day
type DayElapsed time.Duration

// NewDayElapsed creates new DayElapsed from passed time.
func NewDayElapsed(current time.Time, shift time.Duration) DayElapsed {
	current = current.UTC()
	midnight := current.Truncate(Day)
	resetTimeForCurrentDay := midnight.Add(shift)

	if current.Before(resetTimeForCurrentDay) {
		resetTimeForCurrentDay = resetTimeForCurrentDay.Add(-Day) // if we are before reset time, it means reset time was yesterday
	}

	sub := current.Sub(resetTimeForCurrentDay)
	return DayElapsed(sub)
}

// FullHours get full hours of current day
func (e DayElapsed) FullHours() int {
	return int(math.Floor(e.hours()))
}

// HourPart elapsed time of current hour (percentage)
func (e DayElapsed) HourPart() float64 {
	_, f := math.Modf(e.hours())
	return f
}

// Remaining returns remaining time of the day
func (e DayElapsed) Remaining() time.Duration {
	remaining := Day - time.Duration(e)
	if remaining < 0 {
		remaining = 0
	}
	return remaining
}

func (e DayElapsed) hours() float64 {
	return time.Duration(e).Hours()
}
