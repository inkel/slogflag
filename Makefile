checks:
	go vet
	staticcheck

lint:
	golangci-lint run

build: lint test
	go build

test: checks
	go test $(TESTFLAGS)

BENCHFLAGS ?= -bench=. -benchmem

bench:
	go test $(BENCHFLAGS)

watch-test:
	ls -1 *.go | entr -c make test

watch-bench:
	ls -1 *.go | entr -c make bench
