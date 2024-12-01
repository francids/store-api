.PHONY: build run start test

BINARY_NAME=store

install:
	@go mod download

build: install
	@go build -o bin/$(BINARY_NAME) main.go

run: build
	@./bin/$(BINARY_NAME)

start: install
	@go run main.go

test: install
	@go test ./tests/...

clean:
	@rm -rf bin/
