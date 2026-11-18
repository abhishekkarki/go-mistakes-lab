# Usage examples:
#   make test                              all tests, race detector, shuffled order
#   make lab N=026                         run one lab verbosely
#   make bench PKG=lab/026-slices-memory-leaks
#   make benchcmp PKG=kvd/internal/store   compare bench-baseline.txt vs bench.txt
#   make escape PKG=kvd/internal/store     escape analysis (Day 29)
#   make fuzz PKG=kvd/internal/store FUZZ=FuzzParseKey

PKG      ?= kvd/internal/store
N        ?= 000
FUZZ     ?= Fuzz
FUZZTIME ?= 60s

.PHONY: test short integration lab bench baseline benchcmp escape lint cover fuzz run tools

test:
	go test -race -shuffle=on ./...

short:
	go test -short ./...

integration:
	go test -race -shuffle=on -tags integration ./...

lab:
	go test -race -v ./$(wildcard lab/$(N)-*)/...

bench:
	go test -run='^$$' -bench=. -benchmem -count=6 ./$(PKG)/ | tee $(PKG)/bench.txt

baseline:
	go test -run='^$$' -bench=. -benchmem -count=6 ./$(PKG)/ | tee $(PKG)/bench-baseline.txt

benchcmp:
	benchstat $(PKG)/bench-baseline.txt $(PKG)/bench.txt

escape:
	go build -gcflags='-m=2' ./$(PKG)/ 2>&1 | grep -v '^#'

lint:
	golangci-lint run ./...

cover:
	go test -coverprofile=coverage.out ./kvd/...
	go tool cover -html=coverage.out -o coverage.html
	go tool cover -func=coverage.out | tail -1

fuzz:
	go test -run='^$$' -fuzz=$(FUZZ) -fuzztime=$(FUZZTIME) ./$(PKG)/

run:
	go run ./kvd/cmd/kvd

tools:
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest
	go install golang.org/x/perf/cmd/benchstat@latest

