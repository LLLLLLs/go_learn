SHELL := /bin/bash

GO_ENV := GOCACHE=/tmp/go-build CGO_CFLAGS=-Wno-deprecated-declarations
WARN_FILTER := awk 'index($$0, "warning:") == 0 && index($$0, "NSUserNotification") == 0 && index($$0, "duplicate libraries") == 0 && index($$0, "fyne.io/fyne/app") == 0 && index($$0, "stockCalculator.test") == 0 && $$0 !~ / stockCalculator$$/'

.PHONY: tidy build run test

tidy:
	$(GO_ENV) go mod tidy

build:
	@set -o pipefail; $(GO_ENV) go build ./... 2>&1 | $(WARN_FILTER)

run:
	@set -o pipefail; $(GO_ENV) go run . 2>&1 | $(WARN_FILTER)

test:
	@set -o pipefail; $(GO_ENV) go test ./... 2>&1 | $(WARN_FILTER)
