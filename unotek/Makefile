.PHONY: all run build test clean docker-up docker-down swagger migrate-up

APP_NAME=unotek-api

all: build

# Run application locally
run:
	go run ./cmd/api/main.go

# Build production binary
build:
	go build -o bin/$(APP_NAME) ./cmd/api/main.go

# Run tests
test:
	go test -v ./...

# Regenerate Swagger documentation
swagger:
	swag init -g cmd/api/main.go -o docs/

# Start PostgreSQL & API via Docker Compose
docker-up:
	docker compose up -d

# Stop Docker Compose services
docker-down:
	docker compose down

# Stop Docker Compose and remove volumes
docker-clean:
	docker compose down -v

# Clean build artifacts
clean:
	rm -rf bin/
