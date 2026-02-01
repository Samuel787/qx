.PHONY: build install dev clean help

BINARY_NAME=qx
VERSION?=0.1.0
LDFLAGS=-ldflags "-X main.version=$(VERSION)"

help:
	@echo "Available targets:"
	@echo "  build    - Build the binary for macOS"
	@echo "  install  - Build and install the binary locally"
	@echo "  dev      - Run the tool in development mode"
	@echo "  clean    - Remove build artifacts"
	@echo "  deps     - Download and tidy dependencies"

build:
	go build $(LDFLAGS) -o bin/$(BINARY_NAME) ./cmd/qx

install: build
	cp bin/$(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)

dev:
	go run $(LDFLAGS) ./cmd/qx/main.go

clean:
	rm -rf bin/
	go clean

deps:
	go mod download
	go mod tidy

test:
	go test -v ./...
