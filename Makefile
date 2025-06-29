# Matrix - Go Fyne Matrix Effect
# Makefile for building, running, and managing the project

# Variables
BINARY_NAME=matrix
SOURCE_FILE=main.go
BUILD_DIR=build
GO_VERSION=1.21

# Default target
.DEFAULT_GOAL := help

# Build the application
.PHONY: build
build: ## Build the Matrix application
	@echo "Building $(BINARY_NAME)..."
	go build -o $(BINARY_NAME) $(SOURCE_FILE)
	@echo "Build complete: $(BINARY_NAME)"

# Run the application
.PHONY: run
run: ## Run the Matrix application
	@echo "Running $(BINARY_NAME)..."
	go run $(SOURCE_FILE)

# Clean build artifacts
.PHONY: clean
clean: ## Clean build artifacts and binaries
	@echo "Cleaning build artifacts..."
	rm -f $(BINARY_NAME)
	rm -rf $(BUILD_DIR)
	go clean
	@echo "Clean complete"

# Install dependencies
.PHONY: deps
deps: ## Download and install dependencies
	@echo "Installing dependencies..."
	go mod download
	go mod tidy
	@echo "Dependencies installed"

# Build for multiple platforms
.PHONY: build-all
build-all: ## Build for multiple platforms (Linux, macOS, Windows)
	@echo "Building for multiple platforms..."
	mkdir -p $(BUILD_DIR)

	# macOS (Darwin)
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 $(SOURCE_FILE)
	GOOS=darwin GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 $(SOURCE_FILE)

	# Linux
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 $(SOURCE_FILE)
	GOOS=linux GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64 $(SOURCE_FILE)

	# Windows
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe $(SOURCE_FILE)
	GOOS=windows GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-windows-arm64.exe $(SOURCE_FILE)

	@echo "Cross-platform builds complete in $(BUILD_DIR)/"

# Build optimized release version
.PHONY: release
release: ## Build optimized release version
	@echo "Building release version..."
	go build -ldflags="-s -w" -o $(BINARY_NAME) $(SOURCE_FILE)
	@echo "Release build complete: $(BINARY_NAME)"

# Development build with debugging info
.PHONY: dev
dev: ## Build development version with debugging
	@echo "Building development version..."
	go build -gcflags="all=-N -l" -o $(BINARY_NAME) $(SOURCE_FILE)
	@echo "Development build complete: $(BINARY_NAME)"

# Format Go code
.PHONY: fmt
fmt: ## Format Go source code
	@echo "Formatting Go code..."
	go fmt ./...
	@echo "Code formatting complete"

# Lint Go code
.PHONY: lint
lint: ## Lint Go source code
	@echo "Linting Go code..."
	go vet ./...
	@echo "Linting complete"

# Check Go modules
.PHONY: mod-check
mod-check: ## Verify dependencies and check for updates
	@echo "Checking Go modules..."
	go mod verify
	go list -u -m all
	@echo "Module check complete"

# Install the binary to GOPATH/bin
.PHONY: install
install: ## Install the binary to GOPATH/bin
	@echo "Installing $(BINARY_NAME)..."
	go install $(SOURCE_FILE)
	@echo "Installation complete"

# Show project info
.PHONY: info
info: ## Show project information
	@echo "=== Matrix Project Info ==="
	@echo "Binary Name: $(BINARY_NAME)"
	@echo "Source File: $(SOURCE_FILE)"
	@echo "Go Version: $(GO_VERSION)"
	@echo "Current GOOS: $(shell go env GOOS)"
	@echo "Current GOARCH: $(shell go env GOARCH)"
	@echo "=========================="

# Quick build and run
.PHONY: quick
quick: build run ## Quick build and run

# Test build (just compile, don't run)
.PHONY: compile
compile: ## Compile without running (alias for build)
compile: build

# Show available targets
.PHONY: help
help: ## Show this help message
	@echo "Matrix - Go Fyne Matrix Effect"
	@echo "Usage: make [target]"
	@echo ""
	@echo "Available targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-12s %s\n", $$1, $$2}' $(MAKEFILE_LIST)
	@echo ""
	@echo "Examples:"
	@echo "  make build       # Build the application"
	@echo "  make run         # Run the application"
	@echo "  make quick       # Build and run in one command"
	@echo "  make build-all   # Build for all platforms"
	@echo "  make clean       # Clean build artifacts"