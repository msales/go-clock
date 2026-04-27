# Technology Stack

## Language & Runtime
Go 1.22

## Dependencies
- `github.com/benbjohnson/clock v1.3.5` -- real and mock clock implementations
- `github.com/stretchr/testify v1.9.0` -- test assertions and mock framework
- `github.com/stretchr/objx v0.5.2` -- object assertions (indirect, via testify)

## Build System
- `Makefile` includes shared `github.com/msales/make/golang`
- `Dockerfile` uses `msales/go-builder:1.22-base-1.0.1` to verify compilation (`go build -v ./...`)
- No binary output -- library only

## Testing
- `testify/assert` for value assertions
- `testify/mock` for interface mocking
- Custom mockClock implementing `clock.Clock`
- Race condition tests with goroutine coordination
- Benchmark tests (locked vs no-lock performance)
- 3 test files: clock_test.go, clock_bench_test.go, race_test.go

## CI/CD
- GitHub Actions workflow `.github/workflows/test.yaml`
- Runs on ubuntu-latest
- Executes: go test, go vet, staticcheck (with race detection and parallel mode)
- Staticcheck ignores: ST1000, SA1019, ST1003
