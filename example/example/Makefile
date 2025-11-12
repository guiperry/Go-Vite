# example Build System
PROJECT_NAME := example
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_TIME := $(shell date -u +"example-%m-%dT%H:%M:%SZ")
GIT_COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")

ROOT_DIR := $(shell pwd)
BACKEND_DIR := $(ROOT_DIR)/backend
FRONTEND_DIR := $(ROOT_DIR)/frontend
BUILD_DIR := $(ROOT_DIR)/build
DIST_DIR := $(ROOT_DIR)/dist

GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)
CGO_ENABLED := 0
CGO_ENABLED_WEBVIEW := 1

LDFLAGS := -X main.Version=$(VERSION) \
           -X main.BuildTime=$(BUILD_TIME) \
           -X main.GitCommit=$(GIT_COMMIT) \
           -w -s

NPM_CMD := npm

.PHONY: help all clean build test deps frontend backend binary

all: clean deps frontend backend binary

help:
	@echo "Available targets:"
	@echo "  all       - Build everything"
	@echo "  deps      - Install dependencies"
	@echo "  frontend  - Build React frontend"
	@echo "  backend   - Build Go backend"
	@echo "  binary    - Build unified binary"
	@echo "  clean     - Clean build artifacts"
	@echo "  test      - Run tests"

deps: deps-go deps-node

deps-go:
	@echo "Installing Go dependencies..."
	cd $(BACKEND_DIR) && go mod download && go mod tidy
	go mod download

deps-node:
	@echo "Installing Node.js dependencies..."
	cd $(FRONTEND_DIR) && $(NPM_CMD) install

frontend: deps-node
	@echo "Building React frontend..."
	cd $(FRONTEND_DIR) && $(NPM_CMD) run build
	@echo "Frontend build completed"

backend: deps-go
	@echo "Building Go backend..."
	cd $(BACKEND_DIR) && go build -ldflags "$(LDFLAGS)" -o bin/backend ./cmd/server
	@mkdir -p bin
	cp $(BACKEND_DIR)/bin/backend bin/
	@echo "Backend build completed"

binary: frontend backend
	@echo "Creating unified binary..."
	@mkdir -p $(DIST_DIR)
	CGO_ENABLED=$(CGO_ENABLED_WEBVIEW) go build -ldflags "$(LDFLAGS)" -o $(DIST_DIR)/example .
	@echo "Unified binary created: $(DIST_DIR)/example"

clean:
	@echo "Cleaning build artifacts..."
	rm -rf $(BUILD_DIR)
	rm -rf $(DIST_DIR)
	rm -rf $(BACKEND_DIR)/bin
	rm -rf bin

test:
	@echo "Running tests..."
	cd $(BACKEND_DIR) && go test -v ./...

run: binary
	./$(DIST_DIR)/%!s(MISSING)
