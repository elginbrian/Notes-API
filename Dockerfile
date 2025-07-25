# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Generate Swagger documentation
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init

# Build the application
RUN go build -o main .

# Production stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/main .

# Copy the docs directory (Swagger files)
COPY --from=builder /app/docs ./docs

# Create uploads directory
RUN mkdir -p uploads

# Expose port
EXPOSE 3000

# Run the binary
CMD ["./main"]
