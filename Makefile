GOPATH ?= $(shell go env GOPATH)
BIN_DIR := $(GOPATH)/bin
GOLANGCI_LINT := $(BIN_DIR)/golangci-lint
IMAGE_GOLANGCI_LINT ?= golangci/golangci-lint:v1.64.8

.PHONY: docker-lint git-sync go install-vcli lint lint-install lint-test rv release-version test test-html work-http

docker-lint:
	@echo ">> docker lint"
	docker run --rm -v $(shell pwd):/app -w /app $(IMAGE_GOLANGCI_LINT) golangci-lint run -v --timeout=5m
	@echo ">> done"

git-sync:
	@echo ">> git sync"
	git remote update origin -p
	@echo ">> done"

go:
	@echo ">> go"
	go mod tidy && go mod vendor
	@echo ">> done"

install-vcli:
	@echo ">> install vcli"
	cd vcli && go install
	@echo ">> done"

lint:
	@echo ">> run golangci-lint"
	@$(GOLANGCI_LINT) --version
	$(GOLANGCI_LINT) run --timeout=5m
	@echo ">> done"

lint-install:
	@echo ">> install golangci-lint"
	@echo "BIN_DIR:" $(BIN_DIR)
	@echo "GOLANGCI_LINT:" $(GOLANGCI_LINT)
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(BIN_DIR) v1.64.8
	@echo ">> done"

lint-test: lint test

rv: release-version

release-version: work-http go lint-test
	@echo ">> release version"
	make work-http
	cd examples/http && make go && make lint-test && cd ../../
	make work-tcp
	cd examples/tcp && make go && make lint-test && cd ../../
	make work-websocket
	cd examples/websocket && make go && make lint-test && cd ../../
	make work-http
	@echo ">> done"

test:
	@echo ">> run tests"
	go test $$(go list ./... | grep -v "/vendor\|/examples") -count=1 -gcflags="all=-N -l" -parallel=1 -coverprofile=coverage.out
	go tool cover -func coverage.out | tail -n 1 | awk '{ print "Total coverage: " $$3 }'
	@echo ">> done"

test-html: test
	@echo ">> run tests html"
	go tool cover -html coverage.out
	@echo ">> done"

work-http:
	@echo ">> work http"
	cp go.work.http go.work
	@echo ">> done"
