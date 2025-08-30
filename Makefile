# Temporal Orchestration Makefile

# Variables
BINARY_DIR=bin
WORKER_BINARY=$(BINARY_DIR)/worker
CLIENT_BINARY=$(BINARY_DIR)/client

# Go build flags
LDFLAGS=-ldflags "-s -w"

# Default target
.PHONY: all
all: build

# Create binary directory
$(BINARY_DIR):
	mkdir -p $(BINARY_DIR)

# Build all binaries
.PHONY: build
build: $(BINARY_DIR) build-worker build-client

# Build worker
.PHONY: build-worker
build-worker: $(BINARY_DIR)
	go build $(LDFLAGS) -o $(WORKER_BINARY) ./worker

# Build client
.PHONY: build-client
build-client: $(BINARY_DIR)
	go build $(LDFLAGS) -o $(CLIENT_BINARY) ./client

# Run worker
.PHONY: worker
worker: build-worker
	$(WORKER_BINARY)

# Run client
.PHONY: client
client: build-client
	$(CLIENT_BINARY)

# Run both worker and client (in background)
.PHONY: run
run: build
	@echo "Starting worker..."
	@$(WORKER_BINARY) &
	@echo "Worker started with PID: $$!"
	@sleep 2
	@echo "Starting client..."
	@$(CLIENT_BINARY)
	@echo "Stopping worker..."
	@kill $$!

# Run worker in development mode
.PHONY: dev-worker
dev-worker:
	@echo "Starting worker in development mode..."
	@go run ./worker

# Run client in development mode
.PHONY: dev-client
dev-client:
	@echo "Starting client in development mode..."
	@go run ./client

# Run both in development mode
.PHONY: dev
dev:
	@echo "Starting worker in development mode..."
	@go run ./worker &
	@echo "Worker started with PID: $$!"
	@sleep 2
	@echo "Starting client in development mode..."
	@go run ./client
	@echo "Stopping worker..."
	@kill $$!

# Clean build artifacts
.PHONY: clean
clean:
	rm -rf $(BINARY_DIR)
	go clean

# Format code
.PHONY: fmt
fmt:
	go fmt ./...

# Help
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  build         - Build all binaries"
	@echo "  worker        - Run worker"
	@echo "  client        - Run client"
	@echo "  run           - Run both worker and client"
	@echo "  dev-worker    - Run worker in development mode"
	@echo "  dev-client    - Run client in development mode"
	@echo "  dev           - Run both in development mode"
	@echo "  clean         - Clean build artifacts"
	@echo "  fmt           - Format code"
	@echo "  help          - Show this help message"
