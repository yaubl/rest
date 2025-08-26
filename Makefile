APP_NAME := api
CMD_DIR := .
BIN_DIR := bin
GO_FILES := $(shell find . -type f -name '*.go' -not -path "./vendor/*")

GIN_MODE := release
GO := go
GOFLAGS := -mod=readonly
LDFLAGS := -s -w -X main.version=$(VERSION)
BUILDFLAGS := -trimpath -ldflags="$(LDFLAGS)"

VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")

.PHONY: all
all: build

.PHONY: build
build:
	@mkdir -p $(BIN_DIR)
	$(GO) build $(GOFLAGS) $(BUILDFLAGS) -o $(BIN_DIR)/$(APP_NAME) ./$(CMD_DIR)

.PHONY: run
run: build
	env GIN_MODE=$(GIN_MODE) ./$(BIN_DIR)/$(APP_NAME)

.PHONY: test
test:
	$(GO) test ./...

.PHONY: fmt
fmt:
	$(GO) fmt ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: clean
clean:
	rm -rf $(BIN_DIR)

.PHONY: tidy
tidy:
	$(GO) mod tidy

.PHONY: deps
deps:
	$(GO) mod download

