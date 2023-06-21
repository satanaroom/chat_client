.PHONY: build
build:
	rm -rf bin/chat-client
	go build -o bin/chat-client cmd/main.go

.PHONY: lint
lint:
	golangci-lint run