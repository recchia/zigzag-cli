-include .env
export $(shell [ -f .env ] && sed 's/=.*//' .env)

BINARY_NAME=zigzag
MAIN_FILE=cmd/zigzag/main.go
# If ZIGZAG_OUTPUT_DIR is not set, use a sensible default like .
OUTPUT_DIR ?= .
ZIGZAG_OUTPUT_DIR := $(or $(ZIGZAG_OUTPUT_DIR),$(OUTPUT_DIR))

LDFLAGS=-ldflags "-X main.DefaultOutputDir=$(ZIGZAG_OUTPUT_DIR)"

.PHONY: all build test clean

all: test build

test:
	go test ./...

build:
	go build $(LDFLAGS) -o $(BINARY_NAME) $(MAIN_FILE)

clean:
	rm -f $(BINARY_NAME)
