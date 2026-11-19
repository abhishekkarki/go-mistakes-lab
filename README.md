# go-mistakes-lab

Hands-on workbook for *100 Go Mistakes and How to Avoid Them* (Teiva Harsanyi).
Every mistake gets a failing test, benchmark or race report, then a fix.

## Layout

```text
lab/NNN-slug/        one package per mistake (doc.go says which day); write *_test.go here
lab/000-hello/       toolchain smoke test
kvd/                 capstone: in-memory KV server that grows each week
  cmd/kvd/           main
  internal/store/    the store (Day 6 onward)
  internal/api/      HTTP handlers (Day 22)
  internal/client/   Go client (Day 21-22)
NOTES.md             one line per mistake: the rule, then why
```

Every `TODO(day N, #M)` in kvd is a planned change: the workbook day and the mistake it applies.
Search for them with `grep -rn "TODO(day" kvd`.

## Daily loop

1. **Guess**: read the mistake title and write one line in `NOTES.md` saying what you think goes wrong.
2. **Read** the section in the book.
3. **Break it**: write a failing test, benchmark or race in `lab/NNN-*/`.
4. **Fix it** and watch the result change.
5. **Hunt** for the pattern in your own code (inference-router, llm-d).
6. **Write it down**: finish the line in `NOTES.md`.

## Commands

| Command | What it does |
| --- | --- |
| `make tools` | install golangci-lint v2 and benchstat |
| `make test` | all tests, `-race`, shuffled order |
| `make lab N=026` | run one lab, verbose |
| `make bench PKG=lab/026-slices-memory-leaks` | 6-count benchmark, writes `bench.txt` |
| `make baseline PKG=kvd/internal/store` | save `bench-baseline.txt` |
| `make benchcmp PKG=kvd/internal/store` | `benchstat` baseline vs current |
| `make escape PKG=kvd/internal/store` | escape analysis (`-gcflags=-m=2`) |
| `make fuzz PKG=... FUZZ=FuzzX` | fuzz for 60 s |
| `make cover` | kvd coverage report |
| `make integration` | include `//go:build integration` tests |
| `make lint` | golangci-lint |
| `make run` | start kvd on :8080 (`curl localhost:8080/healthz`) |

On Day 11, add goleak with `go get go.uber.org/goleak`. On Day 16, add errgroup with `go get golang.org/x/sync`.

## Go version notes

- `go.mod` says `go 1.24`, so loop variables are per-iteration. To reproduce #32 and #63, use a scratch module with `go 1.21`.
- #76: timers are GC-able since Go 1.23, so measure instead of trusting the book.
- #100: GOMAXPROCS respects cgroup CPU limits since Go 1.25. Compare 1.24 and 1.25 in Docker.
