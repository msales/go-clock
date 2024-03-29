package goclock_test

import (
	"context"
	"testing"
	"time"

	clock2 "github.com/benbjohnson/clock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	goclock "github.com/msales/go-clock/v2"
)

func TestMock(t *testing.T) {
	now := time.Date(2019, time.September, 30, 14, 30, 00, 00, time.UTC)

	mock := goclock.Mock(now)

	time.Sleep(time.Nanosecond) // We just want to be sure that ANY time has passed.
	assert.Equal(t, now, goclock.Now())

	mock.Add(time.Second)
	assert.Equal(t, now.Add(time.Second), goclock.Now())
}

func TestRestore(t *testing.T) {
	now := time.Date(2019, time.September, 30, 14, 30, 00, 00, time.UTC)

	goclock.Mock(now)
	goclock.Restore()

	time.Sleep(time.Nanosecond) // We just want to be sure that ANY time has passed.
	assert.NotEqual(t, now, goclock.Now())
}

func TestAfter(t *testing.T) {
	d := time.Second
	ch := make(<-chan time.Time)

	clk := new(mockClock)
	clk.On("After", d).Return(ch)
	goclock.Set(clk)

	got := goclock.After(d)

	clk.AssertExpectations(t)
	assert.Equal(t, ch, got)
}

func TestAfterFunc(t *testing.T) {
	d := time.Second
	f := func() {}
	timer := &clock2.Timer{}

	clk := new(mockClock)
	clk.On("AfterFunc", d, mock.AnythingOfType("func()")).Return(timer)
	goclock.Set(clk)

	got := goclock.AfterFunc(d, f)

	clk.AssertExpectations(t)
	assert.Equal(t, timer, got)
}

func TestNow(t *testing.T) {
	now := time.Now()

	clk := new(mockClock)
	clk.On("Now").Return(now)
	goclock.Set(clk)

	got := goclock.Now()

	clk.AssertExpectations(t)
	assert.Equal(t, now, got)
}

func TestSince(t *testing.T) {
	now := time.Now()
	d := time.Second

	clk := new(mockClock)
	clk.On("Since", now).Return(d)
	goclock.Set(clk)

	got := goclock.Since(now)

	clk.AssertExpectations(t)
	assert.Equal(t, d, got)
}

func TestSleep(t *testing.T) {
	d := time.Second

	clk := new(mockClock)
	clk.On("Sleep", d)
	goclock.Set(clk)

	goclock.Sleep(d)

	clk.AssertExpectations(t)
}

func TestTick(t *testing.T) {
	d := time.Second
	ch := make(<-chan time.Time)

	clk := new(mockClock)
	clk.On("Tick", d).Return(ch)
	goclock.Set(clk)

	got := goclock.Tick(d)

	clk.AssertExpectations(t)
	assert.Equal(t, ch, got)
}

func TestTicker(t *testing.T) {
	d := time.Second
	ticker := &clock2.Ticker{}

	clk := new(mockClock)
	clk.On("Ticker", d).Return(ticker)
	goclock.Set(clk)

	got := goclock.Ticker(d)

	clk.AssertExpectations(t)
	assert.Equal(t, ticker, got)
}

func TestTimer(t *testing.T) {
	d := time.Second
	timer := &clock2.Timer{}

	clk := new(mockClock)
	clk.On("Timer", d).Return(timer)
	goclock.Set(clk)

	got := goclock.Timer(d)

	clk.AssertExpectations(t)
	assert.Equal(t, timer, got)
}

func TestUntil(t *testing.T) {
	now := time.Date(2019, time.September, 30, 14, 30, 00, 00, time.UTC)
	dur := 5 * time.Second

	clk := new(mockClock)
	clk.On("Until", now).Return(dur)
	goclock.Set(clk)

	got := goclock.Until(now)

	clk.AssertExpectations(t)
	assert.Equal(t, dur, got)
}

func TestWithDeadline(t *testing.T) {
	ctx := context.Background()
	now := time.Now().Add(5 * time.Second)

	clk := new(mockClock)
	clk.On("WithDeadline", ctx, now).Return(ctx, context.CancelFunc(func() {}))
	goclock.Set(clk)

	gotCtx, cancelFn := goclock.WithDeadline(ctx, now)

	clk.AssertExpectations(t)
	assert.Equal(t, ctx, gotCtx)
	assert.IsType(t, context.CancelFunc(func() {}), cancelFn)
}

func TestWithTimeout(t *testing.T) {
	ctx := context.Background()
	timeout := 5 * time.Second

	clk := new(mockClock)
	clk.On("WithTimeout", ctx, timeout).Return(ctx, context.CancelFunc(func() {}))
	goclock.Set(clk)

	gotCtx, cancelFn := goclock.WithTimeout(ctx, timeout)

	clk.AssertExpectations(t)
	assert.Equal(t, ctx, gotCtx)
	assert.IsType(t, context.CancelFunc(func() {}), cancelFn)
}

type mockClock struct {
	mock.Mock
}

func (m *mockClock) Until(t time.Time) time.Duration {
	return m.Called(t).Get(0).(time.Duration)
}

func (m *mockClock) WithDeadline(parent context.Context, d time.Time) (context.Context, context.CancelFunc) {
	return m.Called(parent, d).Get(0).(context.Context), m.Called(parent, d).Get(1).(context.CancelFunc)
}

func (m *mockClock) WithTimeout(parent context.Context, t time.Duration) (context.Context, context.CancelFunc) {
	return m.Called(parent, t).Get(0).(context.Context), m.Called(parent, t).Get(1).(context.CancelFunc)
}

func (m *mockClock) After(d time.Duration) <-chan time.Time {
	return m.Called(d).Get(0).(<-chan time.Time)
}

func (m *mockClock) AfterFunc(d time.Duration, f func()) *clock2.Timer {
	return m.Called(d, f).Get(0).(*clock2.Timer)
}

func (m *mockClock) Now() time.Time {
	return m.Called().Get(0).(time.Time)
}

func (m *mockClock) Since(t time.Time) time.Duration {
	return m.Called(t).Get(0).(time.Duration)
}

func (m *mockClock) Sleep(d time.Duration) {
	m.Called(d)
}

func (m *mockClock) Tick(d time.Duration) <-chan time.Time {
	return m.Called(d).Get(0).(<-chan time.Time)
}

func (m *mockClock) Ticker(d time.Duration) *clock2.Ticker {
	return m.Called(d).Get(0).(*clock2.Ticker)
}

func (m *mockClock) Timer(d time.Duration) *clock2.Timer {
	return m.Called(d).Get(0).(*clock2.Timer)
}
