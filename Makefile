# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOSEC=gosec
GOLINT=golangci-lint

# Main package path
PKG=./

# Binary name
BINARY_NAME=app

# Target build directory
BUILD_DIR=./bin

# Default target
all: test build

# Build the binary
build:
	@echo "Building binary..."
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) $(PKG)

# Clean up build artifacts
clean:
	@echo "Cleaning up..."
	$(GOCLEAN)
	cd $(BUILD_DIR) && del $(BINARY_NAME)

# Run tests
test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

# Run tests with cover
cover:
	@echo "Running tests with cover..."
	$(GOTEST) -cover ./...

# Install dependencies
install:
	@echo "Installing dependencies..."
	$(GOGET) ./...

# Run the application
run:
	@echo "Running the application..."
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) $(PKG)
	$(BUILD_DIR)/$(BINARY_NAME)

# Build and run the application
start: build run

# Create test coverage report
report:
	@echo "Creating test coverage report..."
	$(GOTEST) -covermode=count -coverpkg=./... -coverprofile coverage.out -v ./...
	$(GOCMD) tool cover -html coverage.out -o coverage.html

security-check:
	@echo "Running security check..."
	$(GOSEC) -quiet ./...

vet:
	@echo "Running vet..."
	$(GOCMD) vet ./...

lint:
	@echo "Running lint..."
	$(GOLINT) run ./...

# Help
help:
	@echo "Available targets:"
	@echo "  - build:          Build the binary"
	@echo "  - clean:          Clean up build artifacts"
	@echo "  - test:           Run tests"
	@echo "  - cover:          Run tests with cover"
	@echo "  - report:         Create test coverage report"
	@echo "  - install:        Install dependencies"
	@echo "  - run:            Run the application"
	@echo "  - start:          Build and run the application"
	@echo "  - vet:            Run vet"
	@echo "  - lint:           Run lint"
	@echo "  - security-check: Run security check"
	@echo "  - help:           Display this help message"

# Ensure targets not clash with files in the directory
.PHONY: all build clean test deps run start help
