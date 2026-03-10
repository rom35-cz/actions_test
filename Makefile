.PHONY: all build run test fmt vet clean

BINARY_NAME=actions_test
GO_VERSION=1.24.1

all: fmt vet test build

build:
	go build -o $(BINARY_NAME) .

run: build
	./$(BINARY_NAME)

test:
	go test -v ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

clean:
	rm -f $(BINARY_NAME)
