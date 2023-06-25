SHELL:=/bin/bash


all: build

test:
	go test ./... -v

build:
	go build -o example ./main.go

clean:
	rm -rf example

run:
	go run main.go

.PHONY: all build clean run
