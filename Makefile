# Meba CLI Makefile

.PHONY: build install clean test lint fmt help

# Variables
BINARY_NAME=meba
BUILD_DIR=bin
MAIN_PATH=./main.go

# Default target
help: ## Show this help message
	@echo "Meba CLI - Build Commands"
	@echo "========================="
	@awk 'BEGIN {FS = ":.*##"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

build: ## Build the binary
	@echo "🔨 Building meba CLI..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)
	@echo "✅ Build completed: $(BUILD_DIR)/$(BINARY_NAME)"

install: build ## Install the binary to GOPATH/bin
	@echo "📦 Installing meba CLI..."
	@go install $(MAIN_PATH)
	@echo "✅ Meba CLI installed successfully!"
	@echo "💡 You can now use 'meba' command globally"

clean: ## Clean build artifacts
	@echo "🧹 Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)
	@go clean
	@echo "✅ Clean completed"

test: ## Run tests
	@echo "🧪 Running tests..."
	@go test -v ./...
	@echo "✅ Tests completed"

test-coverage: ## Run tests with coverage
	@echo "🧪 Running tests with coverage..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "✅ Tests completed with coverage report: coverage.html"

lint: ## Run linter
	@echo "🔍 Running linter..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "⚠️  golangci-lint not installed. Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

fmt: ## Format code
	@echo "🎨 Formatting code..."
	@go fmt ./...
	@echo "✅ Code formatted"

deps: ## Download dependencies
	@echo "📥 Downloading dependencies..."
	@go mod download
	@go mod tidy
	@echo "✅ Dependencies updated"

dev-setup: ## Setup development environment
	@echo "🛠️  Setting up development environment..."
	@go mod download
	@go install github.com/cosmtrek/air@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install github.com/swaggo/swag/cmd/swag@latest
	@echo "✅ Development environment ready!"

release: clean test build ## Build release version
	@echo "🚀 Building release..."
	@mkdir -p $(BUILD_DIR)/release
	@GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $(BUILD_DIR)/release/$(BINARY_NAME)-linux-amd64 $(MAIN_PATH)
	@GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o $(BUILD_DIR)/release/$(BINARY_NAME)-darwin-amd64 $(MAIN_PATH)
	@GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o $(BUILD_DIR)/release/$(BINARY_NAME)-darwin-arm64 $(MAIN_PATH)
	@GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o $(BUILD_DIR)/release/$(BINARY_NAME)-windows-amd64.exe $(MAIN_PATH)
	@echo "✅ Release builds completed in $(BUILD_DIR)/release/"

docker-build: ## Build Docker image
	@echo "🐳 Building Docker image..."
	@docker build -t meba-cli:latest .
	@echo "✅ Docker image built: meba-cli:latest"

# Development commands
run-example: ## Run example generation
	@echo "🎯 Running example generation..."
	@mkdir -p tmp/example
	@cd tmp/example && ../../$(BUILD_DIR)/$(BINARY_NAME) new test-app
	@echo "✅ Example project created in tmp/example/test-app"

check: fmt lint test ## Run all checks (format, lint, test)
	@echo "✅ All checks passed!"

.DEFAULT_GOAL := help