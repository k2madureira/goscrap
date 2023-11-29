# Build the aplication
all: build

build:
		@echo "Building..."
		@go build -o main cmd/goscrap/main.go

# Run the aplication
run:
		@go run cmd/goscrap/main.go

# Run the aplication with air
run-air:
		@air cmd/goscrap/main.go

# Test the aplication
test:
		@echo "Testing..."
		@go test ./...

# Clean the binary
clean:
		@echo "Cleaning ..."
		@rm -f main

.PHONY: all build run test clean
