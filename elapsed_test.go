package clock_test

import (
	"testing"
	"time"

	"github.com/msales/go-clock"
	"github.com/stretchr/testify/assert"
)

func TestDayElapsed_FullHours_and_HourPart(t *testing.T) {
	tests := []struct {
		name          string
		dayElapsed    clock.DayElapsed
		wantFullHour  int
		wantHourPart  float64
		wantRemaining time.Duration
	}{
		{
			name:          "midnight",
			dayElapsed:    clock.DayElapsed(time.Duration(0)),
			wantFullHour:  0,
			wantHourPart:  0.0,
			wantRemaining: clock.Day,
		},
		{
			name:          "00:30",
			dayElapsed:    clock.DayElapsed(30 * time.Minute),
			wantFullHour:  0,
			wantHourPart:  0.5,
			wantRemaining: 23*time.Hour + 30*time.Minute,
		},
		{
			name:          "01:00",
			dayElapsed:    clock.DayElapsed(60 * time.Minute),
			wantFullHour:  1,
			wantHourPart:  0.0,
			wantRemaining: 23 * time.Hour,
		},
		{
			name:          "01:30",
			dayElapsed:    clock.DayElapsed(90 * time.Minute),
			wantFullHour:  1,
			wantHourPart:  0.5,
			wantRemaining: 22*time.Hour + 30*time.Minute,
		},
		{
			name:          "22:59",
			dayElapsed:    clock.DayElapsed(1379 * time.Minute),
			wantFullHour:  22,
			wantHourPart:  0.9833333333,
			wantRemaining: 1*time.Hour + 1*time.Minute,
		},
		{
			name:          "23:59",
			dayElapsed:    clock.DayElapsed(1439 * time.Minute),
			wantFullHour:  23,
			wantHourPart:  0.9833333333,
			wantRemaining: 1 * time.Minute,
		},
		{
			name:          "16:37",
			dayElapsed:    clock.DayElapsed(997 * time.Minute),
			wantFullHour:  16,
			wantHourPart:  0.6166666667,
			wantRemaining: 7*time.Hour + 23*time.Minute,
		},
		{
			name:          "25:30 overdue",
			dayElapsed:    clock.DayElapsed(25*time.Hour + 30*time.Minute),
			wantFullHour:  25,
			wantHourPart:  0.5,
			wantRemaining: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFullHours := tt.dayElapsed.FullHours()
			gotHourPart := tt.dayElapsed.HourPart()
			gotRemaining := tt.dayElapsed.Remaining()
			assert.Equal(t, tt.wantFullHour, gotFullHours)
			assert.InDelta(t, tt.wantHourPart, gotHourPart, 0.00001)
			assert.Equal(t, tt.wantRemaining, gotRemaining)
		})
	}
}

func TestDayElapsed_NewElapsed(t *testing.T) {
	tests := []struct {
		name        string
		resetTime   time.Duration
		now         time.Time
		wantElapsed clock.DayElapsed
	}{
		{
			name:        "midnight",
			resetTime:   0,
			now:         time.Date(2020, 12, 30, 12, 0, 0, 0, time.UTC),
			wantElapsed: clock.DayElapsed(12 * time.Hour),
		},
		{
			name:        "3 AM",
			resetTime:   3 * time.Hour,
			now:         time.Date(2020, 12, 30, 12, 0, 0, 0, time.UTC),
			wantElapsed: clock.DayElapsed(9 * time.Hour),
		},
		{
			name:        "15 PM",
			resetTime:   15 * time.Hour,
			now:         time.Date(2020, 12, 30, 20, 0, 0, 0, time.UTC),
			wantElapsed: clock.DayElapsed(5 * time.Hour),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := clock.NewDayElapsed(tt.now, tt.resetTime)
			assert.Equal(t, tt.wantElapsed, got)
		})
	}
}
