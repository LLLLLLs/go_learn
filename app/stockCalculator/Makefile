SHELL := /bin/bash

GO_ENV := GOCACHE=/tmp/go-build CGO_CFLAGS=-Wno-deprecated-declarations
WARN_FILTER := awk 'index($$0, "warning:") == 0 && index($$0, "NSUserNotification") == 0 && index($$0, "duplicate libraries") == 0 && index($$0, "fyne.io/fyne/app") == 0 && index($$0, "stockCalculator.test") == 0 && $$0 !~ / stockCalculator$$/'
DOCKER_PLATFORM ?= linux/amd64
DOCKER_IMAGE ?= stockcalculator-webapp:latest
REMOTE_COMPOSE_DIR ?= /root/stock
SERVER_USER ?= root
SERVER_PORT ?= 22

export SERVER_IP
export SERVER_PASSWORD
export SERVER_USER
export SERVER_PORT
export REMOTE_COMPOSE_DIR

.PHONY: tidy build run run-web test docker-build docker-build-amd64 docker-build-arm64 docker-push docker-deploy

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
	@DEPLOY_AFTER_PUSH=0 ./scripts/deploy-image.sh

docker-deploy: docker-build
	@SKIP_BUILD=1 IMAGE_NAME=$(DOCKER_IMAGE) REMOTE_IMAGE_NAME=$(DOCKER_IMAGE) ./scripts/deploy-image.sh
