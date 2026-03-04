BINARY_NAME=zigzag
MAIN_FILE=cmd/zigzag/main.go

.PHONY: all build test clean

all: test build

test:
	go test ./...

build:
	go build -o $(BINARY_NAME) $(MAIN_FILE)

clean:
	rm -f $(BINARY_NAME)
