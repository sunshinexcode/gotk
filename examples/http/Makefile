GOPATH ?= $(shell go env GOPATH)
BIN_DIR := $(GOPATH)/bin
BRANCH_NAME ?= $(shell git name-rev --name-only HEAD | cut -d / -f 3 | tr / -)
BUILD_DIR := ./build
CMD_DIR := ./cmd
DOCKER_LABELS ?= git-describe="$(shell date -u +v%Y%m%d)-$(shell git describe --tags --always --dirty)"
GENTOOL_DSN := "root:@tcp(localhost:3306)/gotk_example_http?charset=utf8mb4&parseTime=True&loc=Local"
GOLANGCI_LINT := $(BIN_DIR)/golangci-lint
GOLANGCI_LINT_CI := golangci-lint
HELM_DEBUG_PARAM := "--dry-run --debug"
IMAGE_GOLANG ?= golang:1.24.0
IMAGE_GOLANGCI_LINT ?= golangci/golangci-lint:v1.64.8
OUTPUT_DIR := ./bin
PROJECT_NAME := gotk-example-http
PROJECT_NAME_CRON := gotk-example-http-cron
PROJECT_VERSION ?= $(shell echo $(BRANCH_NAME)"_0.1."$(shell date -u +'%Y%m%d%H')-$(shell git rev-parse --short HEAD) | tr / -)
REGISTRY ?=
USER = $(shell whoami)
LDFLAGS := "-s -w                                                                                                         \
           -X github.com/sunshinexcode/gotk/vversion.Branch=`git rev-parse --abbrev-ref HEAD`                     \
           -X github.com/sunshinexcode/gotk/vversion.BuildTime=`date -u '+%Y-%m-%d_%H:%M:%S'`                     \
           -X github.com/sunshinexcode/gotk/vversion.Commit=`git rev-parse HEAD`                                  \
           -X github.com/sunshinexcode/gotk/vversion.Project=$(PROJECT_NAME)                                      \
           -X github.com/sunshinexcode/gotk/vversion.Version=$(PROJECT_VERSION)                                   \
         "

.PHONY: build build-linux build-local clean                                                                     \
	docker-build docker-build-cron docker-build-push docker-build-push-cron docker-clean docker-clean-none      \
	docker-lint docker-push docker-push-cron docker-run docker-run-cron                                         \
	gen-orm git-sync go lint lint-ci lint-install lint-test pprof run run-ab run-cron test test-html

build: build-local

build-linux: clean
	@echo ">> build linux"
	@echo "PWD:"$(PWD)
	docker run --rm -t                                                                                          \
	  -v $(PWD):/go/src/$(PROJECT_NAME)                                                                         \
	  -w /go/src/$(PROJECT_NAME)                                                                                \
	  -e GOOS=linux                                                                                             \
	  -e GOARCH=amd64                                                                                           \
	  -e GOPATH=/go                                                                                             \
	  -e CGO_ENABLED=0                                                                                          \
	  -e GOFLAGS=-mod=vendor                                                                                    \
	  -e SHELLOPTS=errexit                                                                                      \
	  $(IMAGE_GOLANG)                                                                                           \
		/bin/bash -c '                                                                                          \
		   echo "go env" &&                                                                                     \
		   go env &&                                                                                            \
		   echo "go build" &&                                                                                   \
		   go build -ldflags $(LDFLAGS) -buildvcs=false -v -o $(OUTPUT_DIR)/$(PROJECT_NAME) $(CMD_DIR)          \
		   '
	@echo ">> done"

build-local: clean
	@echo ">> build local"
	@GOOS=$(shell uname -s | tr A-Z a-z) GOARCH=$(ARCH) CGO_ENABLED=0 GOFLAGS=-mod=vendor       \
	go env &&                                                                                   \
	go build -ldflags $(LDFLAGS) -v -o $(OUTPUT_DIR)/$(PROJECT_NAME) $(CMD_DIR)
	@echo ">> done"

clean:
	@echo ">> clean up"
	rm -vrf ${OUTPUT_DIR}
	@echo ">> done"

docker-build: build-linux
	@echo ">> docker build"
	@echo "REGISTRY:" $(REGISTRY)
	@echo "PROJECT_NAME:" $(PROJECT_NAME)
	@echo "DOCKER_LABELS:" $(DOCKER_LABELS)
	@echo "BUILD_DIR:" $(BUILD_DIR)
	@echo "PROJECT_VERSION:" $(PROJECT_VERSION)
	docker build -t $(REGISTRY)$(PROJECT_NAME):$(PROJECT_VERSION) --label $(DOCKER_LABELS) -f $(BUILD_DIR)/Dockerfile .
	@echo ">> done"

docker-build-cron: build-linux
	@echo ">> docker build cron"
	@echo "REGISTRY:" $(REGISTRY)
	@echo "PROJECT_NAME_CRON:" $(PROJECT_NAME_CRON)
	@echo "DOCKER_LABELS:" $(DOCKER_LABELS)
	@echo "BUILD_DIR:" $(BUILD_DIR)
	@echo "PROJECT_VERSION:" $(PROJECT_VERSION)
	docker build -t $(REGISTRY)$(PROJECT_NAME_CRON):$(PROJECT_VERSION) --label $(DOCKER_LABELS) -f $(BUILD_DIR)/Dockerfile_cron .
	@echo ">> done"

docker-build-push: docker-build docker-push

docker-build-push-cron: docker-build-cron docker-push-cron

docker-clean: docker-clean-none
	@echo ">> docker clean"
	sh scripts/docker/docker_clean.sh $(REGISTRY)$(PROJECT_NAME)
	@echo ">> done"

docker-clean-none:
	@echo ">> docker clean none"
	sh scripts/docker/docker_clean_none.sh
	@echo ">> done"

docker-lint:
	@echo ">> docker lint"
	docker run --rm -v $(shell pwd):/app -w /app $(IMAGE_GOLANGCI_LINT) golangci-lint run -v --timeout=5m
	@echo ">> done"

docker-push:
	@echo ">> docker push"
	docker push $(REGISTRY)$(PROJECT_NAME):$(PROJECT_VERSION)
	@echo ">> done"

docker-push-cron:
	@echo ">> docker push cron"
	docker push $(REGISTRY)$(PROJECT_NAME_CRON):$(PROJECT_VERSION)
	@echo ">> done"

docker-run:
	@echo ">> docker run"
	docker run --rm -it -p 8080:8080 -e APP_ENV=dev $(REGISTRY)$(PROJECT_NAME):$(PROJECT_VERSION)
	@echo ">> done"

docker-run-cron:
	@echo ">> docker run cron"
	docker run --rm -it -e APP_ENV=dev $(REGISTRY)$(PROJECT_NAME_CRON):$(PROJECT_VERSION)
	@echo ">> done"

gen-orm:
	@echo ">> gen orm"
	gentool -dsn $(GENTOOL_DSN) \
	-tables "" \
	-outPath "internal/entity" \
    -modelPkgName "internal/entity" \
	-onlyModel true
	sed -i '' 's/ID/Id/g' internal/entity/*.gen.go
	@echo ">> done"

git-sync:
	@echo ">> git sync"
	git remote update origin -p
	@echo ">> done"

go:
	@echo ">> go"
	go mod tidy && go mod vendor
	@echo ">> done"

lint:
	@echo ">> run golangci-lint"
	@$(GOLANGCI_LINT) --version
	$(GOLANGCI_LINT) run --timeout=5m
	@echo ">> done"

lint-ci:
	@echo ">> run golangci-lint ci"
	@$(GOLANGCI_LINT_CI) --version
	$(GOLANGCI_LINT_CI) run --timeout=5m
	@echo ">> done"

lint-install:
	@echo ">> install golangci-lint"
	@echo "BIN_DIR:" $(BIN_DIR)
	@echo "GOLANGCI_LINT:" $(GOLANGCI_LINT)
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(BIN_DIR) v1.64.8
	@echo ">> done"

lint-test: lint test

pprof:
	@echo ">> profile"
	@echo ">> p=profile|heap|allocs|goroutine|mutex|block"
	go tool pprof http://localhost:6060/debug/pprof/${p}
	@echo ">> done"

run:
	@echo ">> run"
	APP_ENV=dev go run cmd/main.go http
	@echo ">> done"

run-ab:
	@echo ">> run ab"
	APP_ENV=dev go run cmd/main.go ab $(param)
	@echo ">> done"

run-cron:
	@echo ">> run cron"
	APP_ENV=dev go run cmd/main.go cron
	@echo ">> done"

test:
	@echo ">> run tests"
	go test $$(go list ./... | grep -v /vendor) -count=1 -gcflags="all=-N -l" -parallel=1 -coverprofile=coverage.out
	go tool cover -func coverage.out | tail -n 1 | awk '{ print "Total coverage: " $$3 }'
	@echo ">> done"

test-html: test
	@echo ">> run tests html"
	go tool cover -html coverage.out
	@echo ">> done"
