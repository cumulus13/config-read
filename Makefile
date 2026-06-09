.PHONY: build test clean install lint

APP_NAME = config-read
BUILD_DIR = build
GO_FILES = $(shell find . -name '*.go' -not -path './vendor/*')

# Build variables
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_TIME ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

LDFLAGS = -ldflags "-X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.buildTime=$(BUILD_TIME)"

# Default target
all: clean test build

# Build for current platform
build:
	go build $(LDFLAGS) -o $(BUILD_DIR)/$(APP_NAME) ./cmd/config-read/

# Build for all platforms
build-all:
	@echo "Building for multiple platforms..."
	@mkdir -p $(BUILD_DIR)
	
	# Linux
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(APP_NAME)-linux-amd64 ./cmd/config-read/
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(APP_NAME)-linux-arm64 ./cmd/config-read/
	
	# macOS
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(APP_NAME)-darwin-amd64 ./cmd/config-read/
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(APP_NAME)-darwin-arm64 ./cmd/config-read/
	
	# Windows
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(APP_NAME)-windows-amd64.exe ./cmd/config-read/
	
	@echo "Build complete. Artifacts in $(BUILD_DIR)/"

# Install to GOPATH
install:
	go install $(LDFLAGS) ./cmd/config-read/

# Run tests
test:
	go test -v -race -cover ./...

# Run tests with coverage report
test-coverage:
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Run linter
lint:
	golangci-lint run ./...

# Clean build artifacts
clean:
	rm -rf $(BUILD_DIR)
	rm -f coverage.out coverage.html

# Format code
fmt:
	go fmt ./...

# Run security checks
security:
	govulncheck ./...
	gosec ./...

# Development build with hot reload
dev:
	air -c .air.toml

# Show help
help:
	@echo "Available targets:"
	@echo "  all           : Clean, test, and build"
	@echo "  build         : Build for current platform"
	@echo "  build-all     : Build for all platforms"
	@echo "  install       : Install to GOPATH"
	@echo "  test          : Run tests"
	@echo "  test-coverage : Run tests with coverage report"
	@echo "  lint          : Run linter"
	@echo "  clean         : Clean build artifacts"
	@echo "  fmt           : Format code"
	@echo "  security      : Run security checks"
	@echo "  dev           : Development mode with hot reload"