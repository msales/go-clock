# Concerns & Risks

## Technical Debt
- Global mutable state via package-level `def` singleton -- cannot have multiple independent clock instances in the same process
- `NoLock()` is irreversible (no `EnableLock()` counterpart) -- once disabled, mutex cannot be re-enabled without calling `Restore()`
- `mutexWrap.Disable()` is not thread-safe (writes `disabled` bool without synchronization while other goroutines may be checking it)
- Depends on `benbjohnson/clock` which is archived/unmaintained

## Missing Coverage
- No test for `DayElapsed` edge cases (e.g., exactly at midnight, shift larger than 24h)
- No test for `Current()` function
- No test for `NoLock()` behavior beyond benchmarks
- Race test coordination uses unbuffered-like channel patterns that may not catch all race conditions

## Security
- No security concerns -- pure time library with no I/O or external dependencies

## Performance
- Mutex on every time operation adds overhead; `NoLock()` available as opt-in optimization
- Benchmarks show locked vs no-lock comparison
- `localClock` methods hold mutex during the entire clock operation (could use RWMutex for read-only operations like `Now()`)
