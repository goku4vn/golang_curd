# Makefile for Go CRUD application

# Variables
APP_NAME = crud
BUILD_DIR = ./tmp
PORT ?= 8080

# Default target
.PHONY: all
all: build

# Build the application
.PHONY: build
build:
	go build -o $(BUILD_DIR)/$(APP_NAME) .

# Run tests
.PHONY: test
test:1
	go test ./...

# Clean build artifacts
.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)

# Start development server with Air for live reloading
.PHONY: dev
dev:
	PORT=$(PORT) air -c .air.toml

# Help target
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  all    - Build the application (default)"
	@echo "  build  - Build the application"
	@echo "  test   - Run tests"
	@echo "  clean  - Remove build artifacts"
	@echo "  dev    - Start development server with Air for live reloading (PORT=8080 by default)"
	@echo "  help   - Show this help message"
	@echo ""
	@echo "Variables:"
	@echo "  PORT   - Port for the development server (default: 8080)"
	@echo "          Example: make dev PORT=3001"
