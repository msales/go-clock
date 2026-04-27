# Architecture

## Overview
A thread-safe clock abstraction providing package-level time functions that can be mocked for testing. Wraps benbjohnson/clock with a mutex-protected global singleton and adds day-elapsed calculations.

## Package Structure
- `goclock` (single package) -- entire library in the root directory
  - `clock.go` -- core clock wrapper (localClock), package-level functions, mutex wrapper
  - `elapsed.go` -- DayElapsed type for time-of-day calculations with shift support

## Key Interfaces
- `clock.Clock` (from benbjohnson/clock) -- the underlying interface implemented by both real and mock clocks
- Package-level functions (`Now()`, `After()`, `Sleep()`, `Mock()`, `Restore()`, etc.) -- the primary API surface

## Data Flow
1. Consumer calls package-level function (e.g., `goclock.Now()`)
2. Function delegates to global `def` singleton (`*localClock`)
3. `localClock` acquires mutex (unless disabled via `NoLock()`)
4. Delegates to underlying `clock.Clock` implementation (real or mock)
5. For testing: `Mock(time)` replaces the underlying clock with a frozen `clock.Mock`
6. `Restore()` reverts to real clock

## External Dependencies
None -- pure library with no external service dependencies.
