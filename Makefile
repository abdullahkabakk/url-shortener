# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

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

# Help
help:
	@echo "Available targets:"
	@echo "  - build:     Build the binary"
	@echo "  - clean:     Clean up build artifacts"
	@echo "  - test:      Run tests"
	@echo "  - install:   Install dependencies"
	@echo "  - run:       Run the application"
	@echo "  - start:     Build and run the application"
	@echo "  - help:      Display this help message"

# Ensure targets not clash with files in the directory
.PHONY: all build clean test deps run start help
