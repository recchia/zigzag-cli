BINARY_NAME=zigzag
BUILD_DIR=build
MAIN_FILE=cmd/zigzag/main.go

.PHONY: all build test clean

all: test build

test:
	go test ./...

build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_FILE)

clean:
	rm -rf $(BUILD_DIR)
