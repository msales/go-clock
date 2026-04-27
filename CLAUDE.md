@~/CLAUDE.md

## Project Description

**Language:** Go 1.22  
**Module:** `github.com/msales/go-clock/v2`  
**Purpose:** A thread-safe clock abstraction library that wraps `benbjohnson/clock`. Provides package-level functions for time operations (`Now`, `After`, `Sleep`, `Tick`, etc.) with a global mockable clock instance. Supports mocking for tests, optional mutex locking, and day-elapsed time calculations.  
**Key dependencies:**
- `github.com/benbjohnson/clock` -- underlying clock implementation (real and mock)
- `github.com/stretchr/testify` -- test assertions and mocks

## Directory Structure

```
(root)              Single-package library (package goclock)
.github/workflows/  GitHub Actions CI (test.yaml)
```

Root-level files: `clock.go` (main clock wrapper with mutex), `elapsed.go` (DayElapsed type), test files (`clock_test.go`, `clock_bench_test.go`, `race_test.go`).

## Build, Test & Lint

```bash
# Build
go build ./...

# Run tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests with race detection
go test -race ./...

# Run benchmarks
go test -bench=. ./...

# Vet
go vet ./...

# Static analysis (CI uses staticcheck)
staticcheck ./...
```

Makefile includes `github.com/msales/make/golang` (shared Makefile). Dockerfile builds and compiles the library as a verification step.

## Code Style

### Imports

Three groups: stdlib, external (`benbjohnson/clock`), internal (separated by blank lines).

### Naming Conventions

- **Packages:** `goclock` (Go convention for library packages)
- **Interfaces:** uses `clock.Clock` from the underlying library
- **Exported/unexported:** exported package-level functions (`Now`, `Mock`, `Set`, `Restore`, `NoLock`); unexported `localClock` struct wrapping the real implementation

### Error Handling

Minimal -- time operations do not typically return errors. No error wrapping needed.

### Testing

- `testify/assert` for assertions
- `testify/mock` for mocking the `clock.Clock` interface
- Custom `mockClock` struct implementing the full `clock.Clock` interface
- Dedicated race condition tests (`race_test.go`) with goroutine coordination
- Benchmark tests comparing locked vs no-lock performance

## Branching & Git

- **Never** work directly on `main` or `master`. Always create a new branch off the latest `main`/`master`.
- Before creating a branch, **always pull** the latest changes from the remote (`git pull origin main`).
- If the user provides a task identifier in the format `TRK-XXXX`, use it as the branch name (e.g., `TRK-1234`). Otherwise, **ask the user** for a task ID or branch name before proceeding.
- **Never push** to `main` or `master` directly.
- Commit changes after each meaningful iteration. Use your judgment to decide when progress should be saved -- prefer smaller, atomic commits over large monolithic ones.
- Write clear, concise commit messages. Use conventional commit format when appropriate (e.g., `feat:`, `fix:`, `refactor:`, `test:`).
