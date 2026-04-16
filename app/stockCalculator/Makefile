GO_ENV := GOCACHE=/tmp/go-build

.PHONY: tidy build run

tidy:
	$(GO_ENV) go mod tidy

build:
	$(GO_ENV) go build ./...

run:
	$(GO_ENV) go run .
