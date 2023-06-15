.PHONY: build
build:
	go build -o bin/client cmd/client/main.go

.PHONY: lint
lint:
	golangci-lint run