build:
	go build -o bin ./...

start: build
	./bin

phony: build start
