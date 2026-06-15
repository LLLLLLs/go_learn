SHELL := /bin/bash

GO_ENV := GOCACHE=/tmp/go-build CGO_CFLAGS=-Wno-deprecated-declarations
WARN_FILTER := awk 'index($$0, "warning:") == 0 && index($$0, "NSUserNotification") == 0 && index($$0, "duplicate libraries") == 0 && index($$0, "fyne.io/fyne/app") == 0 && index($$0, "stockCalculator.test") == 0 && $$0 !~ / stockCalculator$$/'
DOCKER_PLATFORM ?= linux/amd64
DOCKER_IMAGE ?= stockcalculator-webapp:latest

.PHONY: tidy build run run-web test docker-build docker-build-amd64 docker-build-arm64 docker-push

tidy:
	$(GO_ENV) go mod tidy

build:
	@set -o pipefail; $(GO_ENV) go build ./... 2>&1 | $(WARN_FILTER)

run:
	@set -o pipefail; $(GO_ENV) go run ./cmd/desktop 2>&1 | $(WARN_FILTER)

run-web:
	@set -o pipefail; $(GO_ENV) go run ./webapp 2>&1 | $(WARN_FILTER)

test:
	@set -o pipefail; $(GO_ENV) go test ./... 2>&1 | $(WARN_FILTER)

docker-build:
	docker build --platform $(DOCKER_PLATFORM) -t $(DOCKER_IMAGE) .

docker-build-amd64:
	$(MAKE) docker-build DOCKER_PLATFORM=linux/amd64 DOCKER_IMAGE=stockcalculator-webapp:amd64

docker-build-arm64:
	$(MAKE) docker-build DOCKER_PLATFORM=linux/arm64 DOCKER_IMAGE=stockcalculator-webapp:arm64

docker-push:
	./scripts/deploy-image.sh
