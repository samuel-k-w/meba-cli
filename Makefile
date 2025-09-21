VERSION ?= $(shell git describe --tags --always --dirty)
LDFLAGS = -ldflags "-X github.com/meba-cli/meba/cmd.version=$(VERSION)"

.PHONY: build
build:
	@echo "🔨 Building meba CLI..."
	@go build $(LDFLAGS) -o bin/meba ./cmd/meba
	@echo ""
	@echo "✅ Build completed: bin/meba"

.PHONY: build-all
build-all:
	@echo "🔨 Building for all platforms..."
	@mkdir -p dist
	@GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o dist/meba-linux-amd64 ./cmd/meba
	@GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o dist/meba-darwin-amd64 ./cmd/meba
	@GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o dist/meba-darwin-arm64 ./cmd/meba
	@GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o dist/meba-windows-amd64.exe ./cmd/meba
	@echo "✅ All binaries built in dist/"

.PHONY: install
install: build
	@echo "📦 Installing meba CLI..."
	@sudo cp bin/meba /usr/local/bin/
	@echo "✅ Installed to /usr/local/bin/meba"

.PHONY: test
test:
	@go test ./...

.PHONY: clean
clean:
	@rm -rf bin/ dist/

.PHONY: release
release:
	@echo "🚀 Creating release..."
	@git tag -a v$(VERSION) -m "Release v$(VERSION)"
	@git push origin v$(VERSION)