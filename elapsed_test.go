package goclock_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	goclock "github.com/msales/go-clock/v2"
)

func TestDayElapsed_FullHours_and_HourPart(t *testing.T) {
	tests := []struct {
		name          string
		dayElapsed    goclock.DayElapsed
		wantFullHour  int
		wantHourPart  float64
		wantRemaining time.Duration
	}{
		{
			name:          "midnight",
			dayElapsed:    goclock.DayElapsed(time.Duration(0)),
			wantFullHour:  0,
			wantHourPart:  0.0,
			wantRemaining: goclock.Day,
		},
		{
			name:          "00:30",
			dayElapsed:    goclock.DayElapsed(30 * time.Minute),
			wantFullHour:  0,
			wantHourPart:  0.5,
			wantRemaining: 23*time.Hour + 30*time.Minute,
		},
		{
			name:          "01:00",
			dayElapsed:    goclock.DayElapsed(60 * time.Minute),
			wantFullHour:  1,
			wantHourPart:  0.0,
			wantRemaining: 23 * time.Hour,
		},
		{
			name:          "01:30",
			dayElapsed:    goclock.DayElapsed(90 * time.Minute),
			wantFullHour:  1,
			wantHourPart:  0.5,
			wantRemaining: 22*time.Hour + 30*time.Minute,
		},
		{
			name:          "22:59",
			dayElapsed:    goclock.DayElapsed(1379 * time.Minute),
			wantFullHour:  22,
			wantHourPart:  0.9833333333,
			wantRemaining: 1*time.Hour + 1*time.Minute,
		},
		{
			name:          "23:59",
			dayElapsed:    goclock.DayElapsed(1439 * time.Minute),
			wantFullHour:  23,
			wantHourPart:  0.9833333333,
			wantRemaining: 1 * time.Minute,
		},
		{
			name:          "16:37",
			dayElapsed:    goclock.DayElapsed(997 * time.Minute),
			wantFullHour:  16,
			wantHourPart:  0.6166666667,
			wantRemaining: 7*time.Hour + 23*time.Minute,
		},
		{
			name:          "25:30 overdue",
			dayElapsed:    goclock.DayElapsed(25*time.Hour + 30*time.Minute),
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

func TestNewDayElapsed(t *testing.T) {
	tests := []struct {
		name          string
		current       time.Time
		shift         time.Duration
		wantElapsed   goclock.DayElapsed
		wantRemaining time.Duration
		wantFullHours int
		wantHourPart  float64
	}{
		{
			name:          "Reset time on midnight",
			current:       time.Date(2021, 3, 15, 15, 59, 30, 0, time.UTC),
			shift:         0, // UTC midnight
			wantElapsed:   goclock.DayElapsed(15*time.Hour + 59*time.Minute + 30*time.Second),
			wantRemaining: 8*time.Hour + 30*time.Second,
			wantFullHours: 15,
			wantHourPart:  0.9916666667,
		},
		{
			name:          "Reset time on midnight at midnight",
			current:       time.Date(2021, 3, 15, 0, 0, 0, 0, time.UTC),
			shift:         0, // UTC midnight
			wantElapsed:   goclock.DayElapsed(0),
			wantRemaining: goclock.Day,
			wantFullHours: 0,
			wantHourPart:  0,
		},
		{
			name:          "Reset time at 13:20, current time at 13:20",
			current:       time.Date(2021, 3, 15, 13, 20, 0, 0, time.UTC),
			shift:         13*time.Hour + 20*time.Minute,
			wantElapsed:   goclock.DayElapsed(0),
			wantRemaining: goclock.Day,
			wantFullHours: 0,
			wantHourPart:  0,
		},
		{
			name:          "Reset time at 16:00 UTC, current at 15:59:30",
			current:       time.Date(2021, 3, 15, 15, 59, 30, 0, time.UTC),
			shift:         16 * time.Hour,
			wantElapsed:   goclock.DayElapsed(23*time.Hour + 59*time.Minute + 30*time.Second),
			wantRemaining: 30 * time.Second,
			wantFullHours: 23,
			wantHourPart:  0.9916666667,
		},
		{
			name:          "Reset time at 15:00 UTC, current at 15:59:30",
			current:       time.Date(2021, 3, 15, 15, 59, 30, 0, time.UTC),
			shift:         15 * time.Hour,
			wantElapsed:   goclock.DayElapsed(59*time.Minute + 30*time.Second),
			wantRemaining: 23*time.Hour + 30*time.Second,
			wantFullHours: 0,
			wantHourPart:  0.9916666667,
		},
		{
			name:          "Reset time at 16:00 UTC, current at 15:59:30 UTC+1",
			current:       time.Date(2021, 3, 15, 15, 59, 30, 0, time.FixedZone("UTC1", 3600)),
			shift:         16 * time.Hour,
			wantElapsed:   goclock.DayElapsed(22*time.Hour + 59*time.Minute + 30*time.Second),
			wantRemaining: 1*time.Hour + 30*time.Second,
			wantFullHours: 22,
			wantHourPart:  0.9916666667,
		},
		{
			name:          "Reset time at 23:00 UTC, current at 00:00:00 UTC+1",
			current:       time.Date(2021, 3, 15, 0, 0, 0, 0, time.FixedZone("UTC1", 3600)),
			shift:         23 * time.Hour,
			wantElapsed:   goclock.DayElapsed(0),
			wantRemaining: goclock.Day,
			wantFullHours: 0,
			wantHourPart:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			elapsed := goclock.NewDayElapsed(tt.current, tt.shift)
			assert.Equal(t, tt.wantElapsed, elapsed)
			assert.Equal(t, tt.wantRemaining, elapsed.Remaining())
			assert.Equal(t, tt.wantFullHours, elapsed.FullHours())
			assert.InDelta(t, tt.wantHourPart, elapsed.HourPart(), 0.0001)
		})
	}
}
