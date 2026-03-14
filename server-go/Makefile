# Variables
BINARY_NAME=server
GO=go
GOFLAGS=-v

.PHONY: tidy
tidy:
	$(GO) mod tidy

.PHONY: build
build: tidy
	CGO_ENABLED=0 $(GO) build $(GOFLAGS) -o $(BINARY_NAME) ./cmd/main.go

.PHONY: test
test: tidy
	$(GO) test $(GOFLAGS) ./...

.PHONY: test-cover
test-cover:
	$(GO) test -v -cover -coverprofile=coverage.out ./...
	$(GO) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

.PHONY: golangci-lint
golangci-lint:
	@command -v golangci-lint > /dev/null || (echo "golangci-lint not found. Install with: https://golangci-lint.run/docs/welcome/install/local/" && exit 1)

.PHONY: lint
lint: golangci-lint
	$(GO) fmt ./...
	$(GO) vet ./...
	golangci-lint run ./...

.PHONY: lint-fix
lint-fix: golangci-lint
	golangci-lint run --fix ./...

.PHONY: run
run: build
	./$(BINARY_NAME)

.PHONY: watch
watch:
	@command -v entr > /dev/null || (echo "entr not found. Install with: brew install entr" && exit 1)
	find . -name "*.go" | entr -r make run

.PHONY: clean
clean:
	$(GO) clean
	rm -f $(BINARY_NAME)
	rm -f coverage.out coverage.html
	@echo "Cleaned build artifacts"

