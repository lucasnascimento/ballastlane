# Output name of binary
BINARY_NAME=clock

# Build the project
build:
	docker compose build

# Clean the project
clean:
	go clean
	rm -f ./bin/$(BINARY_NAME)

# Test the project
test:
	docker compose up postgres -d
	go test -v ./...
	docker compose down -v

# Run the project
run:
	docker compose down -v
	docker compose up

# Install dependencies
deps:
	go mod download

.PHONY: build clean test run deps
