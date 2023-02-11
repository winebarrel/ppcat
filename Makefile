all: vet test build

.PHONY: build
build:
	go build ./cmd/ppcat

.PHONY: vet
vet:
	go vet ./...

.PHONY: test
test:
	go test -v ./...

lint:
	golangci-lint run
