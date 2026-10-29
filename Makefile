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

# Run database schema migrations
migrate-up:
	go run ./cmd/migrate/main.go up

# Run database seeder (initial super admin & master data)
migrate-seed:
	go run ./cmd/migrate/main.go seed

# Run both migration and seeding
migrate-all:
	go run ./cmd/migrate/main.go all

# Drop all tables (reset database)
migrate-drop:
	go run ./cmd/migrate/main.go drop

# Clean build artifacts
clean:
	rm -rf bin/
