# Notes API Makefile

# Build the application
build:
	go build -o notes-api .

# Run the application locally
run:
	go run main.go

# Run with Docker Compose
docker-up:
	docker-compose up --build

# Stop Docker Compose
docker-down:
	docker-compose down

# Run Docker Compose in background
docker-up-bg:
	docker-compose up --build -d

# View logs
docker-logs:
	docker-compose logs -f

# Clean up Docker
docker-clean:
	docker-compose down -v
	docker system prune -f

# Run tests
test:
	go test ./...

# Install dependencies
deps:
	go mod tidy
	go mod download

# Generate Swagger documentation
swagger:
	swag init

# Format code
fmt:
	go fmt ./...

# Run linter
lint:
	golint ./...

# Clean build artifacts
clean:
	rm -f notes-api
	rm -f notes-api.exe

# Create uploads directory
create-uploads:
	mkdir -p uploads

# Database migrate (when running locally)
migrate:
	@echo "Database migration will run automatically when the application starts"

.PHONY: build run docker-up docker-down docker-up-bg docker-logs docker-clean test deps swagger fmt lint clean create-uploads migrate
