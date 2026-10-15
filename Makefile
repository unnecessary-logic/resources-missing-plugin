BINARY  := kubectl-resources_missing
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)

.PHONY: build install clean

build:
	go build -ldflags "-s -w -X main.version=$(VERSION)" -o bin/$(BINARY) .

# Copy onto PATH so `kubectl resources-missing` is discovered.
install: build
	install -d $(shell go env GOPATH)/bin
	install -m 0755 bin/$(BINARY) $(shell go env GOPATH)/bin/$(BINARY)

clean:
	rm -rf bin

test:
	go test ./...